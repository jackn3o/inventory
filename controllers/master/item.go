package master

import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	"../../base/connector"
	utility "../../base/utilities"
)

// Controller is return value for New method
type Controller struct {
	databaseName string
	store        *connector.Store
	config       configuration.Config
}

// New method is a constructor for controller.
func New(store *connector.Store, config configuration.Config) *Controller {
	return &Controller{
		databaseName: config.GetString(configuration.DatabaseName),
		store:        store,
		config:       config,
	}
}

const (
	ItemCollection = "item"
)

type Item struct {
	Code           string    `bson:"code" json:"code"`
	Description    string    `bson:"description" json:"description"`
	Color          string    `bson:"color" json:"color"`
	Category       string    `bson:"category" json:"category"`
	OpeningBalance int       `bson:"openingBal" json:"openingBal"`
	CreatedDate    time.Time `bson:"createdDate" json:"createdDate"`
	CreatedBy      string    `bson:"createBy" json:"createBy"`
	ModifiedDate   time.Time `bson:"modifiedDate" json:"modifiedDate"`
	ModifiedBy     string    `bson:"modifiedBy" json:"modifiedBy"`
}

// CreateItem in master data
func (c *Controller) CreateItem() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		session := c.store.DB

		// Collection
		collection := session.DB(c.databaseName).C(ItemCollection)

		// Insert
		if err := collection.Insert(Item{
			Description: "test",
		}); err != nil {
			panic(err)
		}
		u.WriteJSON("OK")
	})
}
