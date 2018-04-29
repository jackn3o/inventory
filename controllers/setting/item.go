package setting

import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	"../../base/connector"
	utility "../../base/utilities"
	"github.com/gorilla/mux"
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
	ID             bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Code           string        `bson:"code" json:"code" valid:"required"`
	Description    string        `bson:"description" json:"description" valid:"required"`
	Color          string        `bson:"color,omitempty" json:"color" valid:"-"`
	Category       string        `bson:"category" json:"category" valid:"-"`
	OpeningBalance int           `bson:"openingBalance" json:"openingBalance,string" valid:"required"`
	CreatedDate    time.Time     `bson:"createdDate" json:"createdDate" valid:"-"`
	CreatedBy      string        `bson:"createBy" json:"createBy" valid:"-"`
	ModifiedDate   time.Time     `bson:"modifiedDate,omitempty" json:"modifiedDate" valid:"-"`
	ModifiedBy     string        `bson:"modifiedBy,omitempty" json:"modifiedBy" valid:"-"`
}

// CreateItem in master data
func (c *Controller) CreateItem() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		item := &Item{}
		err := u.UnmarshalWithValidation(item)
		if err != nil {
			u.WriteJSONError(err.Error(), http.StatusBadRequest)
			return
		}
		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemsCollection)
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

// GetItems return list of items
func (c *Controller) GetItems() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		u.WriteJSON("success")
	})
}

// GetItemByID ...
func (c *Controller) GetItemByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		itemID := mux.Vars(req)["id"]

		if len(itemID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		u.WriteJSON(itemID)
	})
}
