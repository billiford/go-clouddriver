package kubernetes

type Resource struct {
	AccountName string `json:"accountName"`
	ID          string `json:"id" gorm:"primary_key"`
	TaskID      string `json:"taskId"`
	APIGroup    string `json:"apiGroup"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Resource    string `json:"resource"`
	Version     string `json:"version"`
	Kind        string `json:"kind"`
}

func (Resource) TableName() string {
	return "resource_kubernetes"
}
