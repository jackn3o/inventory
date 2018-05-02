package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Item Model for settings.items Collection
type Item struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Code         string        `bson:"code" json:"code" valid:"required"`
	Description  string        `bson:"description" json:"description" valid:"required"`
	Balance      int           `bson:"balance" json:"balance,string" valid:"-"`
	Cost         int           `bson:"cost" json:"cost" valid:"-"`
	Color        string        `bson:"color,omitempty" json:"color" valid:"-"`
	Category     string        `bson:"category,omitempty" json:"category" valid:"-"`
	Outlet       string        `bson:"outlet,omitempty" json:"outlet" valid:"-"`
	CreatedDate  *time.Time    `bson:"createdDate" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// CreateItem in master data
func (c *Controller) CreateItem() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		item := &Item{}
		err := u.UnmarshalWithValidation(item)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}
		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemSettingCollection)
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
		timeNow := time.Now()
		item.CreatedDate = &timeNow
		item.CreatedBy = "todo"
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
		session := c.store.DB.Copy()
		defer session.Close()

		var items []Item
		collection := session.DB(c.databaseName).C(ItemSettingCollection)
		err := collection.Find(nil).Sort("description").All(&items)
		if err != nil {
			u.WriteJSONError("verification failed", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(items)
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

		var item Item
		collection := session.DB(c.databaseName).C(ItemSettingCollection)
		err := collection.FindId(bson.ObjectIdHex(itemID)).One(&item)
		if err != nil {
			u.WriteJSONError("verification failed", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(item)
	})
}
