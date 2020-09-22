package core

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	clouddriver "github.com/billiford/go-clouddriver/pkg"
	"github.com/billiford/go-clouddriver/pkg/arcade"
	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type Applications []Application

type Application struct {
	Attributes   ApplicationAttributes `json:"attributes"`
	ClusterNames map[string][]string   `json:"clusterNames"`
	Name         string                `json:"name"`
}

type ApplicationAttributes struct {
	Name string `json:"name"`
}

func ListApplications(c *gin.Context) {
	sc := sql.Instance(c)

	rs, err := sc.ListKubernetesResourcesByFields("account_name", "kind", "name", "spinnaker_app")
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	response := Applications{}
	apps := uniqueSpinnakerApps(rs)

	for _, app := range apps {
		application := Application{
			Attributes: ApplicationAttributes{
				Name: app,
			},
			ClusterNames: clusterNamesForSpinnakerApp(app, rs),
			Name:         app,
		}

		response = append(response, application)
	}

	c.JSON(http.StatusOK, response)
}

func uniqueSpinnakerApps(rs []kubernetes.Resource) []string {
	apps := []string{}

	for _, r := range rs {
		if !contains(apps, r.SpinnakerApp) {
			apps = append(apps, r.SpinnakerApp)
		}
	}

	return apps
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func clusterNamesForSpinnakerApp(application string, rs []kubernetes.Resource) map[string][]string {
	clusterNames := map[string][]string{}

	for _, r := range rs {
		if r.SpinnakerApp == application {
			if _, ok := clusterNames[r.AccountName]; !ok {
				clusterNames[r.AccountName] = []string{}
			}
			resources := clusterNames[r.AccountName]
			resources = append(resources, fmt.Sprintf("%s %s", r.Kind, r.Name))
			clusterNames[r.AccountName] = resources
		}
	}

	return clusterNames
}

type ServerGroupManagers []ServerGroupManager

type ServerGroupManager struct {
	Account       string                          `json:"account"`
	AccountName   string                          `json:"accountName"`
	CloudProvider string                          `json:"cloudProvider"`
	CreatedTime   int64                           `json:"createdTime"`
	Key           Key                             `json:"key"`
	Kind          string                          `json:"kind"`
	Labels        map[string]string               `json:"labels"`
	Manifest      map[string]interface{}          `json:"manifest"`
	Moniker       Moniker                         `json:"moniker"`
	Name          string                          `json:"name"`
	ProviderType  string                          `json:"providerType"`
	Region        string                          `json:"region"`
	ServerGroups  []ServerGroupManagerServerGroup `json:"serverGroups"`
	Type          string                          `json:"type"`
	UID           string                          `json:"uid"`
	Zone          string                          `json:"zone"`
}

type Key struct {
	Account        string `json:"account"`
	Group          string `json:"group"`
	KubernetesKind string `json:"kubernetesKind"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	Provider       string `json:"provider"`
}

type Moniker struct {
	App     string `json:"app"`
	Cluster string `json:"cluster"`
}

type ServerGroupManagerServerGroup struct {
	Account   string                               `json:"account"`
	Moniker   ServerGroupManagerServerGroupMoniker `json:"moniker"`
	Name      string                               `json:"name"`
	Namespace string                               `json:"namespace"`
	Region    string                               `json:"region"`
}

type ServerGroupManagerServerGroupMoniker struct {
	App      string `json:"app"`
	Cluster  string `json:"cluster"`
	Sequence int    `json:"sequence"`
}

// Server Group Managers for a kubernetes target are deployments.
func ListServerGroupManagers(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.ControllerInstance(c)
	ac := arcade.Instance(c)
	application := c.Param("application")
	response := ServerGroupManagers{}

	accounts, err := sc.ListKubernetesAccountsBySpinnakerApp(application)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Don't actually return while attempting to create a list of server group managers.
	// We want to avoid the situation where a user cannot perform operations when any
	// cluster is not available.
	for _, account := range accounts {
		provider, err := sc.GetKubernetesProvider(account)
		if err != nil {
			log.Println("unable to get kubernetes provider for account", account)
			continue
		}

		cd, err := base64.StdEncoding.DecodeString(provider.CAData)
		if err != nil {
			log.Println("error decoding ca data for account", account)
			continue
		}

		token, err := ac.Token()
		if err != nil {
			log.Println("error getting token", err.Error())
			continue
		}

		config := &rest.Config{
			Host:        provider.Host,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: cd,
			},
		}

		client, err := kc.NewClient(config)
		if err != nil {
			log.Println("error creating dynamic client for account", account)
			continue
		}

		deployments := &unstructured.UnstructuredList{}
		replicaSets := &unstructured.UnstructuredList{}

		lo := metav1.ListOptions{
			LabelSelector: kubernetes.LabelKubernetesName + "=" + application,
		}

		deploymentGVR := schema.GroupVersionResource{
			Group:    "apps",
			Version:  "v1",
			Resource: "deployments",
		}
		replicaSetGVR := schema.GroupVersionResource{
			Group:    "apps",
			Version:  "v1",
			Resource: "replicasets",
		}

		deployments, err = client.ListByGVR(deploymentGVR, lo)
		if err != nil {
			log.Println("error listing deployments:", err.Error())
			continue
		}

		replicaSets, err = client.ListByGVR(replicaSetGVR, lo)
		if err != nil {
			log.Println("error listing replicaSets:", err.Error())
			continue
		}

		for _, deployment := range deployments.Items {
			sgm := newServerGroupManager(deployment, account, application)
			sgm.ServerGroups = buildServerGroups(replicaSets, deployment, account, application)
			response = append(response, sgm)
		}
	}

	c.JSON(http.StatusOK, response)
}

func newServerGroupManager(deployment unstructured.Unstructured,
	account, application string) ServerGroupManager {
	return ServerGroupManager{
		Account:       account,
		AccountName:   account,
		CloudProvider: "kubernetes",
		CreatedTime:   deployment.GetCreationTimestamp().Unix() * 1000,
		Key: Key{
			Account:        account,
			Group:          "deployment",
			KubernetesKind: "deployment",
			Name:           deployment.GetName(),
			Namespace:      deployment.GetNamespace(),
			Provider:       "kubernetes",
		},
		Kind:     "deployment",
		Labels:   deployment.GetLabels(),
		Manifest: deployment.Object,
		Moniker: Moniker{
			App:     application,
			Cluster: fmt.Sprintf("%s %s", "deployment", deployment.GetName()),
		},
		Name:         fmt.Sprintf("%s %s", "deployment", deployment.GetName()),
		ProviderType: "kubernetes",
		Region:       deployment.GetNamespace(),
		Type:         "kubernetes",
		UID:          string(deployment.GetUID()),
		Zone:         application,
	}
}

func buildServerGroups(replicaSets *unstructured.UnstructuredList,
	deployment unstructured.Unstructured,
	account, application string) []ServerGroupManagerServerGroup {
	sgs := []ServerGroupManagerServerGroup{}

	// Deployments manage replicasets, so build a list of managed replicasets for each deployment.
	for _, replicaSet := range replicaSets.Items {
		annotations := replicaSet.GetAnnotations()
		if annotations != nil {
			name := annotations["artifact.spinnaker.io/name"]
			t := annotations["artifact.spinnaker.io/type"]
			if strings.EqualFold(name, deployment.GetName()) &&
				strings.EqualFold(t, "kubernetes/deployment") {
				sequence, _ := strconv.Atoi(annotations["deployment.kubernetes.io/revision"])
				s := ServerGroupManagerServerGroup{
					Account: account,
					Moniker: ServerGroupManagerServerGroupMoniker{
						App:      application,
						Cluster:  fmt.Sprintf("%s %s", "deployment", deployment.GetName()),
						Sequence: sequence,
					},
					Name:      fmt.Sprintf("%s %s", "replicaSet", replicaSet.GetName()),
					Namespace: replicaSet.GetNamespace(),
					Region:    replicaSet.GetNamespace(),
				}
				sgs = append(sgs, s)
			}
		}
	}

	return sgs
}

type LoadBalancers []LoadBalancer

type LoadBalancer struct {
	Account       string                    `json:"account"`
	CloudProvider string                    `json:"cloudProvider"`
	DispatchRules []interface{}             `json:"dispatchRules,omitempty"`
	HTTPURL       string                    `json:"httpUrl,omitempty"`
	HTTPSURL      string                    `json:"httpsUrl,omitempty"`
	Labels        map[string]string         `json:"labels,omitempty"`
	Moniker       Moniker                   `json:"moniker"`
	Name          string                    `json:"name"`
	Project       string                    `json:"project,omitempty"`
	Region        string                    `json:"region"`
	SelfLink      string                    `json:"selfLink,omitempty"`
	ServerGroups  []LoadBalancerServerGroup `json:"serverGroups"`
	Type          string                    `json:"type"`
	AccountName   string                    `json:"accountName,omitempty"`
	CreatedTime   int64                     `json:"createdTime,omitempty"`
	Key           Key                       `json:"key,omitempty"`
	Kind          string                    `json:"kind,omitempty"`
	Manifest      map[string]interface{}    `json:"manifest,omitempty"`
	ProviderType  string                    `json:"providerType,omitempty"`
	UID           string                    `json:"uid,omitempty"`
	Zone          string                    `json:"zone,omitempty"`
}

type LoadBalancerServerGroup struct {
	AllowsGradualTrafficMigration bool          `json:"allowsGradualTrafficMigration"`
	CloudProvider                 string        `json:"cloudProvider"`
	DetachedInstances             []interface{} `json:"detachedInstances"`
	Instances                     []interface{} `json:"instances"`
	IsDisabled                    bool          `json:"isDisabled"`
	Name                          string        `json:"name"`
	Region                        string        `json:"region"`
}

// List "load balancers", which for kubernetes are kinds "ingress" and "service".
func ListLoadBalancers(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.ControllerInstance(c)
	ac := arcade.Instance(c)
	application := c.Param("application")
	response := LoadBalancers{}

	accounts, err := sc.ListKubernetesAccountsBySpinnakerApp(application)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Don't actually return while attempting to create a list of load balancers.
	// We want to avoid the situation where a user cannot perform operations when any
	// cluster is not available.
	for _, account := range accounts {
		provider, err := sc.GetKubernetesProvider(account)
		if err != nil {
			log.Println("unable to get kubernetes provider for account", account)
			continue
		}

		cd, err := base64.StdEncoding.DecodeString(provider.CAData)
		if err != nil {
			log.Println("error decoding ca data for account", account)
			continue
		}

		token, err := ac.Token()
		if err != nil {
			log.Println("error getting token", err.Error())
			continue
		}

		config := &rest.Config{
			Host:        provider.Host,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: cd,
			},
		}

		client, err := kc.NewClient(config)
		if err != nil {
			log.Println("error creating dynamic client for account", account)
			continue
		}

		// Label selector for all that we are listing in the cluster. We
		// only want to list resources that have a label referencing the requested application.
		lo := metav1.ListOptions{
			LabelSelector: kubernetes.LabelKubernetesName + "=" + application,
		}

		// TODO get these using the dynamic account.
		// Create a GVR for ingresses.
		ingressGVR := schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1beta1",
			Resource: "ingresses",
		}

		ingresses, err := client.ListByGVR(ingressGVR, lo)
		if err != nil {
			log.Println("error listing ingresses:", err.Error())
			continue
		}

		for _, ingress := range ingresses.Items {
			lb := newLoadBalancer(ingress, account, application)
			response = append(response, lb)
		}

		// Create a GVR for services.
		serviceGVR := schema.GroupVersionResource{
			Version:  "v1",
			Resource: "services",
		}

		services, err := client.ListByGVR(serviceGVR, lo)
		if err != nil {
			log.Println("error listing services:", err.Error())
			continue
		}

		for _, service := range services.Items {
			lb := newLoadBalancer(service, account, application)
			response = append(response, lb)
		}
	}

	c.JSON(http.StatusOK, response)
}

func newLoadBalancer(u unstructured.Unstructured, account, application string) LoadBalancer {
	kind := strings.ToLower(u.GetKind())
	return LoadBalancer{
		Account:       account,
		AccountName:   account,
		CloudProvider: "kubernetes",
		Labels:        u.GetLabels(),
		Moniker: Moniker{
			App:     application,
			Cluster: fmt.Sprintf("%s %s", kind, u.GetName()),
		},
		Name:        fmt.Sprintf("%s %s", kind, u.GetName()),
		Region:      u.GetNamespace(),
		Type:        "kubernetes",
		CreatedTime: u.GetCreationTimestamp().Unix() * 1000,
		Key: Key{
			Account:        account,
			Group:          u.GroupVersionKind().Group,
			KubernetesKind: kind,
			Name:           fmt.Sprintf("%s %s", kind, u.GetName()),
			Namespace:      u.GetNamespace(),
			Provider:       "kubernetes",
		},
		Kind:         kind,
		Manifest:     u.Object,
		ProviderType: "kubernetes",
		UID:          string(u.GetUID()),
		Zone:         application,
	}
}

type Clusters map[string][]string

// List clusters, which for kubernetes is a map of provider names to kubernetes deployment
// kinds and names.
//
// TODO For now we are pulling this from the DB, but it might be possible to make API calls to
// the cluster.
func ListClusters(c *gin.Context) {
	sc := sql.Instance(c)

	rs, err := sc.ListKubernetesResourcesByFields("account_name", "kind", "name")
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	response := Clusters{}

	for _, resource := range rs {
		if _, ok := response[resource.AccountName]; !ok {
			response[resource.AccountName] = []string{}
		}
		kr := response[resource.AccountName]
		kr = append(kr, fmt.Sprintf("%s %s", resource.Kind, resource.Name))
		response[resource.AccountName] = kr
	}

	c.JSON(http.StatusOK, response)
}

type ServerGroups []ServerGroup

type ServerGroup struct {
	Account        string            `json:"account"`
	AccountName    string            `json:"accountName"`
	BuildInfo      BuildInfo         `json:"buildInfo"`
	Capacity       Capacity          `json:"capacity"`
	CloudProvider  string            `json:"cloudProvider"`
	Cluster        string            `json:"cluster,omitempty"`
	CreatedTime    int64             `json:"createdTime"`
	Disabled       bool              `json:"disabled"`
	InstanceCounts InstanceCounts    `json:"instanceCounts"`
	Instances      []Instance        `json:"instances"`
	IsDisabled     bool              `json:"isDisabled"`
	Key            Key               `json:"key"`
	Kind           string            `json:"kind"`
	Labels         map[string]string `json:"labels"`
	// LaunchConfig struct {} `json:"launchConfig"`
	LoadBalancers       []interface{}                   `json:"loadBalancers"`
	Manifest            map[string]interface{}          `json:"manifest"`
	Moniker             ServerGroupMoniker              `json:"moniker"`
	Name                string                          `json:"name"`
	ProviderType        string                          `json:"providerType"`
	Region              string                          `json:"region"`
	SecurityGroups      []interface{}                   `json:"securityGroups"`
	ServerGroupManagers []ServerGroupServerGroupManager `json:"serverGroupManagers"`
	Type                string                          `json:"type"`
	UID                 string                          `json:"uid"`
	Zone                string                          `json:"zone"`
	Zones               []interface{}                   `json:"zones"`
	InsightActions      []interface{}                   `json:"insightActions"`
}

type ServerGroupServerGroupManager struct {
	Account  string `json:"account"`
	Location string `json:"location"`
	Name     string `json:"name"`
}

type ServerGroupMoniker struct {
	App      string `json:"app"`
	Cluster  string `json:"cluster"`
	Sequence int    `json:"sequence"`
}

type BuildInfo struct {
	Images []string `json:"images"`
}

type Capacity struct {
	Desired int  `json:"desired"`
	Pinned  bool `json:"pinned"`
}

type InstanceCounts struct {
	Down         int `json:"down"`
	OutOfService int `json:"outOfService"`
	Starting     int `json:"starting"`
	Total        int `json:"total"`
	Unknown      int `json:"unknown"`
	Up           int `json:"up"`
}

type Instance struct {
	Account           string                 `json:"account,omitempty"`
	AccountName       string                 `json:"accountName,omitempty"`
	AvailabilityZone  string                 `json:"availabilityZone,omitempty"`
	CloudProvider     string                 `json:"cloudProvider,omitempty"`
	CreatedTime       int64                  `json:"createdTime,omitempty"`
	Health            []InstanceHealth       `json:"health,omitempty"`
	HealthState       string                 `json:"healthState,omitempty"`
	HumanReadableName string                 `json:"humanReadableName,omitempty"`
	ID                string                 `json:"id,omitempty"`
	Key               Key                    `json:"key,omitempty"`
	Kind              string                 `json:"kind,omitempty"`
	Labels            map[string]string      `json:"labels,omitempty"`
	Manifest          map[string]interface{} `json:"manifest,omitempty"`
	Moniker           Moniker                `json:"moniker,omitempty"`
	Name              string                 `json:"name,omitempty"`
	ProviderType      string                 `json:"providerType,omitempty"`
	Region            string                 `json:"region,omitempty"`
	Type              string                 `json:"type,omitempty"`
	UID               string                 `json:"uid,omitempty"`
	Zone              string                 `json:"zone,omitempty"`
}

type InstanceHealth struct {
	Platform string `json:"platform,omitempty"`
	Source   string `json:"source,omitempty"`
	State    string `json:"state"`
	Type     string `json:"type"`
}

func ListServerGroups(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.ControllerInstance(c)
	ac := arcade.Instance(c)
	application := c.Param("application")
	response := ServerGroups{}

	accounts, err := sc.ListKubernetesAccountsBySpinnakerApp(application)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Don't actually return while attempting to create a list of server groups.
	// We want to avoid the situation where a user cannot perform operations when any
	// cluster is not available.
	for _, account := range accounts {
		provider, err := sc.GetKubernetesProvider(account)
		if err != nil {
			log.Println("unable to get kubernetes provider for account", account)
			continue
		}

		cd, err := base64.StdEncoding.DecodeString(provider.CAData)
		if err != nil {
			log.Println("error decoding ca data for account", account)
			continue
		}

		token, err := ac.Token()
		if err != nil {
			log.Println("error getting token", err.Error())
			continue
		}

		config := &rest.Config{
			Host:        provider.Host,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: cd,
			},
		}

		client, err := kc.NewClient(config)
		if err != nil {
			log.Println("error creating dynamic client for account", account)
			continue
		}

		lo := metav1.ListOptions{
			LabelSelector: kubernetes.LabelKubernetesName + "=" + application,
		}

		// Create a GVR for replicasets.
		replicaSetGVR := schema.GroupVersionResource{
			Group:    "apps",
			Version:  "v1",
			Resource: "replicasets",
		}
		podsGVR := schema.GroupVersionResource{
			Version:  "v1",
			Resource: "pods",
		}

		replicaSets, err := client.ListByGVR(replicaSetGVR, lo)
		if err != nil {
			log.Println("error listing replicaSets:", err.Error())
			continue
		}

		pods, err := client.ListByGVR(podsGVR, lo)
		if err != nil {
			log.Println("error listing pods:", err.Error())
			continue
		}

		for _, replicaSet := range replicaSets.Items {
			sg := newServerGroup(replicaSet, pods, account)
			response = append(response, sg)
		}
	}

	c.JSON(http.StatusOK, response)
}

func newServerGroup(replicaSet unstructured.Unstructured,
	pods *unstructured.UnstructuredList, account string) ServerGroup {
	rs := kubernetes.NewReplicaSet(replicaSet.Object)
	images := rs.ListImages()
	spec := rs.GetReplicaSetSpec()

	desired := 0
	if spec.Replicas != nil {
		desired = int(*spec.Replicas)
	}

	serverGroupManagers := []ServerGroupServerGroupManager{}
	instances := buildInstances(pods, replicaSet)
	annotations := replicaSet.GetAnnotations()

	// Build server group manager
	managerName := annotations["artifact.spinnaker.io/name"]
	managerLocation := annotations["artifact.spinnaker.io/location"]
	managerType := annotations["artifact.spinnaker.io/type"]
	if managerType == "kubernetes/deployment" {
		sgm := ServerGroupServerGroupManager{
			Account:  account,
			Location: managerLocation,
			Name:     managerName,
		}
		serverGroupManagers = append(serverGroupManagers, sgm)
	}

	cluster := annotations["moniker.spinnaker.io/cluster"]
	app := annotations["moniker.spinnaker.io/application"]
	sequence, _ := strconv.Atoi(annotations["deployment.kubernetes.io/revision"])

	return ServerGroup{
		Account: account,
		BuildInfo: BuildInfo{
			Images: images,
		},
		Capacity: Capacity{
			Desired: desired,
			Pinned:  false,
		},
		CloudProvider: "kubernetes",
		Cluster:       cluster,
		CreatedTime:   replicaSet.GetCreationTimestamp().Unix() * 1000,
		InstanceCounts: InstanceCounts{
			Down:         0,
			OutOfService: 0,
			Starting:     0,
			Total:        int(rs.GetReplicaSetStatus().Replicas),
			Unknown:      0,
			Up:           int(rs.GetReplicaSetStatus().ReadyReplicas),
		},
		Instances:     instances,
		IsDisabled:    false,
		LoadBalancers: nil,
		Moniker: ServerGroupMoniker{
			App:      app,
			Cluster:  cluster,
			Sequence: sequence,
		},
		Name:                fmt.Sprintf("%s %s", "replicaset", replicaSet.GetName()),
		Region:              replicaSet.GetNamespace(),
		SecurityGroups:      nil,
		ServerGroupManagers: serverGroupManagers,
		Type:                "kubernetes",
		Labels:              replicaSet.GetLabels(),
	}
}

func buildInstances(pods *unstructured.UnstructuredList,
	replicaSet unstructured.Unstructured) []Instance {
	instances := []Instance{}
	for _, u := range pods.Items {
		p := kubernetes.NewPod(u.Object)
		for _, ownerReference := range p.GetObjectMeta().OwnerReferences {
			if strings.EqualFold(ownerReference.Name, replicaSet.GetName()) {
				state := "Up"
				if p.GetPodStatus().Phase != "Running" {
					state = "Down"
				}
				instance := Instance{
					AvailabilityZone: u.GetNamespace(),
					Health: []InstanceHealth{
						{
							State: state,
							Type:  "kubernetes/pod",
						},
						{
							State: state,
							Type:  "kubernetes/container",
						},
					},
					HealthState: state,
					ID:          string(u.GetUID()),
					Name:        fmt.Sprintf("%s %s", "pod", u.GetName()),
				}
				instances = append(instances, instance)
			}
		}
	}
	return instances
}

// /applications/:application/serverGroups/:account/:location/:name
func GetServerGroup(c *gin.Context) {
	sc := sql.Instance(c)
	kc := kubernetes.ControllerInstance(c)
	ac := arcade.Instance(c)
	account := c.Param("account")
	application := c.Param("application")
	location := c.Param("location")
	n := c.Param("name")
	a := strings.Split(n, " ")
	kind := a[0]
	name := a[1]

	provider, err := sc.GetKubernetesProvider(account)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	cd, err := base64.StdEncoding.DecodeString(provider.CAData)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	token, err := ac.Token()
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	config := &rest.Config{
		Host:        provider.Host,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: cd,
		},
	}

	client, err := kc.NewClient(config)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	lo := metav1.ListOptions{
		LabelSelector: kubernetes.LabelKubernetesName + "=" + application,
	}

	podsGVR := schema.GroupVersionResource{
		Version:  "v1",
		Resource: "pods",
	}

	result, err := client.Get(kind, name, location)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// "Instances" in kubernetes are pods.
	pods, err := client.ListByGVR(podsGVR, lo)
	if err != nil {
		clouddriver.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	images := []string{}
	desired := 0
	instanceCounts := InstanceCounts{}
	if strings.EqualFold(kind, "replicaset") {
		rs := kubernetes.NewReplicaSet(result.Object)
		spec := rs.GetReplicaSetSpec()
		status := rs.GetReplicaSetStatus()
		images = rs.ListImages()
		if spec.Replicas != nil {
			desired = int(*spec.Replicas)
		}
		instanceCounts.Total = int(status.Replicas)
		instanceCounts.Up = int(status.ReadyReplicas)
	}

	instances := []Instance{}
	for _, v := range pods.Items {
		p := kubernetes.NewPod(v.Object)
		for _, ownerReference := range p.GetObjectMeta().OwnerReferences {
			if strings.EqualFold(ownerReference.Name, result.GetName()) {
				state := "Up"
				if p.GetPodStatus().Phase != "Running" {
					state = "Down"
				}

				annotations := p.GetObjectMeta().Annotations
				cluster := annotations["moniker.spinnaker.io/cluster"]
				app := annotations["moniker.spinnaker.io/application"]

				if app == "" {
					app = application
				}

				instance := Instance{
					Account:          account,
					AccountName:      account,
					AvailabilityZone: p.GetNamespace(),
					CloudProvider:    "kubernetes",
					CreatedTime:      p.GetObjectMeta().CreationTimestamp.Unix() * 1000,
					Health: []InstanceHealth{
						{
							State: state,
							Type:  "kubernetes/pod",
						},
						{
							State: state,
							Type:  "kubernetes/container",
						},
					},
					HealthState:       state,
					HumanReadableName: fmt.Sprintf("%s %s", "pod", p.GetName()),
					ID:                string(p.GetUID()),
					Key: Key{
						Account:        account,
						Group:          "pod",
						KubernetesKind: "pod",
						Name:           p.GetName(),
						Namespace:      p.GetNamespace(),
						Provider:       "kubernetes",
					},
					Kind:     "pod",
					Labels:   p.GetLabels(),
					Manifest: v.Object,
					Moniker: Moniker{
						App:     app,
						Cluster: cluster,
					},
					Name:         fmt.Sprintf("%s %s", "pod", p.GetName()),
					ProviderType: "kubernetes",
					Region:       p.GetNamespace(),
					Type:         "kubernetes",
					UID:          string(p.GetUID()),
					Zone:         p.GetNamespace(),
				}
				instances = append(instances, instance)
			}
		}
	}

	annotations := result.GetAnnotations()
	cluster := annotations["moniker.spinnaker.io/cluster"]
	app := annotations["moniker.spinnaker.io/application"]
	sequence, _ := strconv.Atoi(annotations["deployment.kubernetes.io/revision"])

	if app == "" {
		app = application
	}

	response := ServerGroup{
		Account:     account,
		AccountName: account,
		BuildInfo: BuildInfo{
			Images: images,
		},
		Capacity: Capacity{
			Desired: desired,
			Pinned:  false,
		},
		CloudProvider:  "kubernetes",
		CreatedTime:    result.GetCreationTimestamp().Unix() * 1000,
		Disabled:       false,
		InstanceCounts: instanceCounts,
		Instances:      instances,
		Key: Key{
			Account:        account,
			Group:          result.GetKind(),
			KubernetesKind: result.GetKind(),
			Name:           result.GetName(),
			Namespace:      result.GetNamespace(),
			Provider:       "kubernetes",
		},
		Kind:          result.GetKind(),
		Labels:        result.GetLabels(),
		LoadBalancers: []interface{}{},
		Manifest:      result.Object,
		Moniker: ServerGroupMoniker{
			App:      app,
			Cluster:  cluster,
			Sequence: sequence,
		},
		Name:                fmt.Sprintf("%s %s", result.GetKind(), result.GetName()),
		ProviderType:        "kubernetes",
		Region:              result.GetNamespace(),
		SecurityGroups:      []interface{}{},
		ServerGroupManagers: []ServerGroupServerGroupManager{},
		Type:                "kubernetes",
		UID:                 string(result.GetUID()),
		Zone:                result.GetNamespace(),
		Zones:               []interface{}{},
		InsightActions:      []interface{}{},
	}

	c.JSON(http.StatusOK, response)
}
