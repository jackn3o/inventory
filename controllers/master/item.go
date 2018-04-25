package master

import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	"../../base/connector"
	utility "../../base/utilities"
	"gopkg.in/mgo.v2/bson"
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

// Key for collection
const (
	ItemsCollection = "items"
)

// Item Model for items Collection
type Item struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Code           string        `bson:"code" json:"code"`
	Description    string        `bson:"description" json:"description"`
	Color          string        `bson:"color,omitempty" json:"color"`
	Category       string        `bson:"category" json:"category"`
	OpeningBalance int           `bson:"openingBal" json:"openingBal"`
	CreatedDate    time.Time     `bson:"createdDate" json:"createdDate"`
	CreatedBy      string        `bson:"createBy" json:"createBy"`
	ModifiedDate   time.Time     `bson:"modifiedDate,omitempty" json:"modifiedDate"`
	ModifiedBy     string        `bson:"modifiedBy,omitempty" json:"modifiedBy"`
}

// CreateItem in master data
func (c *Controller) CreateItem() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemsCollection)
		item := Item{
			Code:           "t",
			Description:    "test",
			Color:          "white",
			Category:       "c",
			OpeningBalance: 1,
			CreatedDate:    time.Now(),
			CreatedBy:      "admin",
		}

		count, err := collection.Find(bson.M{"code": item.Code}).Count()
		if err != nil {
			u.WriteJSONError("verification failed", http.StatusInternalServerError)
			return
		}

		if count != 0 {
			u.WriteJSONError("code exist", http.StatusConflict)
			return
		}

		item.ID = bson.NewObjectId()
		err = collection.Insert(item)

		if err != nil {
			u.WriteJSONError("code exist", http.StatusInternalServerError)
			panic(err)
		}

		u.WriteJSON(item.ID, http.StatusCreated)
	})
}
