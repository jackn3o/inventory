package items

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"gopkg.in/mgo.v2/bson"
)

// Item Model for items Collection
type Item struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Code            string        `bson:"code" json:"code" valid:"required"`
	Description     string        `bson:"description" json:"description" valid:"required"`
	Category        string        `bson:"category,omitempty" json:"category" valid:"-"`
	Color           string        `bson:"color,omitempty" json:"color" valid:"-"`
	UnitCost        float64       `bson:"unitCost" json:"unitCost" valid:"-"`
	BalanceQuantity int           `bson:"balanceQuantity" json:"balanceQuantity" valid:"-"`
	BalanceCost     float64       `bson:"balanceCost" json:"balanceCost" valid:"-"`
	Details         []*ItemDetail `bson:"details,omitempty" json:"details,omitempty" valid:"-"`
	CreatedDate     *time.Time    `bson:"createdDate" json:"createdDate,omitempty" valid:"-"`
	CreatedBy       string        `bson:"createBy" json:"createBy,omitempty" valid:"-"`
	ModifiedDate    *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy      string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
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
		collection := session.DB(c.databaseName).C(ItemsCollection)
		err := collection.
			Find(nil).
			Select(bson.M{"_id": 1, "color": 1, "code": 1, "description": 1, "category": 1}).
			Sort("description").
			All(&items)
		if err != nil {
			u.WriteJSONError("verification failed", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(items)
	})
}
