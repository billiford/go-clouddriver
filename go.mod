module github.com/billiford/go-clouddriver

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/mcuadros/go-gin-prometheus v0.1.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200814200057-3d37ad5750ed // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v11.0.0+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.18.6
