package sql_test

import (
	"database/sql"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	. "github.com/billiford/go-clouddriver/pkg/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sql", func() {
	var (
		db   *gorm.DB
		mock sqlmock.Sqlmock
		d    *sql.DB
		c    Client
		err  error
	)

	BeforeEach(func() {
		// Mock DB.
		d, mock, _ = sqlmock.New()
		db, err = Connect("sqlite3", d)
		// Enable DB logging.
		// db.LogMode(true)
		c = NewClient(db)
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("#Connect", func() {
		When("it fails to connect", func() {
			BeforeEach(func() {
				_, err = Connect("mysql", "mysql")
			})

			It("returns an error", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid DSN: missing the slash separating the database name"))
			})
		})
	})

	Describe("#CreateKubernetesProvider", func() {
		var provider kubernetes.Provider

		BeforeEach(func() {
			provider = kubernetes.Provider{
				Name:   "test-name",
				Host:   "test-host",
				CAData: "test-ca-data",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateKubernetesProvider(provider)
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "kubernetes_providers" \(` +
					`"name",` +
					`"host",` +
					`"ca_data",` +
					`"bearer_token"` +
					`\) VALUES \(\?,\?,\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#GetKubernetesProvider", func() {
		var provider kubernetes.Provider

		BeforeEach(func() {
			provider = kubernetes.Provider{}
		})

		JustBeforeEach(func() {
			provider, err = c.GetKubernetesProvider("test-name")
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"name", "host", "ca_data"}).
					AddRow("test-name", "test-host", "test-ca-data")
				mock.ExpectQuery(`(?i)^SELECT host, ca_data, bearer_token FROM "kubernetes_providers" ` +
					` WHERE \(name = \?\) ORDER BY "kubernetes_providers"."name" ASC LIMIT 1$`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(provider.Name).To(Equal("test-name"))
				Expect(provider.Host).To(Equal("test-host"))
				Expect(provider.CAData).To(Equal("test-ca-data"))
			})
		})
	})

	Describe("#CreateKubernetesResource", func() {
		var resource kubernetes.Resource

		BeforeEach(func() {
			resource = kubernetes.Resource{
				ID:        "test-id",
				TaskID:    "test-task-id",
				APIGroup:  "test-group",
				Name:      "test-name",
				Namespace: "test-namespace",
				Resource:  "test-resource",
				Version:   "test-version",
			}
		})

		JustBeforeEach(func() {
			err = c.CreateKubernetesResource(resource)
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				mock.ExpectBegin()
				mock.ExpectExec(`(?i)^INSERT INTO "kubernetes_resources" \(` +
					`"account_name",` +
					`"id",` +
					`"task_id",` +
					`"api_group",` +
					`"name",` +
					`"namespace",` +
					`"resource",` +
					`"version",` +
					`"kind",` +
					`"spinnaker_app"` +
					`\) VALUES \(\?,\?,\?,\?,\?,\?,\?,\?,\?,\?\)$`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#ListKubernetesResourcesByTaskID", func() {
		var resources []kubernetes.Resource

		JustBeforeEach(func() {
			resources, err = c.ListKubernetesResourcesByTaskID("test-task-id")
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"group", "name"}).
					AddRow("group1", "name1").
					AddRow("group2", "name2")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`account_name, ` +
					`api_group, ` +
					`kind, ` +
					`name, ` +
					`namespace, ` +
					`resource, ` +
					`version ` +
					`FROM "kubernetes_resources" ` +
					` WHERE \(task_id = \?\)$`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(resources).To(HaveLen(2))
			})
		})
	})

	Describe("#ListKubernetesAccountsBySpinnakerApp", func() {
		var accounts []string

		JustBeforeEach(func() {
			accounts, err = c.ListKubernetesAccountsBySpinnakerApp("test-spinnaker-app")
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				sqlRows := sqlmock.NewRows([]string{"account_name"}).
					AddRow("account1").
					AddRow("account2")
				mock.ExpectQuery(`(?i)^SELECT ` +
					`account_name ` +
					`FROM "kubernetes_resources" ` +
					` WHERE \(spinnaker_app = \?\) ` +
					`GROUP BY account_name$`).
					WillReturnRows(sqlRows)
				mock.ExpectCommit()
			})

			It("succeeds", func() {
				Expect(err).To(BeNil())
				Expect(accounts).To(HaveLen(2))
			})
		})
	})

	Describe("#Instance", func() {
		var ctx *gin.Context
		var c2 Client

		BeforeEach(func() {
			ctx = &gin.Context{}
			ctx.Set(ClientInstanceKey, c)
		})

		When("it succeeds", func() {
			BeforeEach(func() {
				c2 = Instance(ctx)
			})

			It("succeeds", func() {
				Expect(c2).ToNot(BeNil())
			})
		})
	})

	Describe("#Connection", func() {
		var connection string
		var c Config

		When("the config is not set", func() {
			BeforeEach(func() {
				connection = Connection(c)
			})

			It("returns a disk db", func() {
				Expect(connection).To(Equal("clouddriver.db"))
			})
		})

		When("the config is set", func() {
			BeforeEach(func() {
				c = Config{
					User:     "user",
					Password: "password",
					Host:     "10.1.1.1",
					Name:     "go-clouddriver",
				}
				connection = Connection(c)
			})

			It("returns a disk db", func() {
				Expect(connection).To(Equal("user:password@tcp(10.1.1.1)/go-clouddriver?charset=utf8&parseTime=True&loc=UTC"))
			})
		})
	})
})
