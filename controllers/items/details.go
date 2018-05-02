package items

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// ItemDetail model for items.details Collection
type ItemDetail struct {
	ID              bson.ObjectId `bson:"id,omitempty" json:"id,omitempty" valid:"-"`
	DocumentNo      string        `bson:"documentNo,omitempty" json:"documentNo,omitempty" valid:"-"`
	BatchNo         string        `bson:"batchNo,omitempty" json:"batchNo,omitempty" valid:"-"`
	Date            *time.Time    `bson:"date,omitempty" json:"date,omitempty" valid:"-"`
	UnitCost        float64       `bson:"unitCost,omitempty" json:"unitCost,string,omitempty" valid:"-"`
	In              float32       `bson:"in,omitempty" json:"in,string,omitempty" valid:"-"`
	Out             float32       `bson:"out,omitempty" json:"out,string,omitempty" valid:"-"`
	BalanceQuantity float32       `bson:"balanceQuantity,omitempty" json:"balanceQuantity,omitempty" valid:"-"`
	CreatedDate     *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy       string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate    *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy      string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// DetailResponseDto model
type DetailResponseDto struct {
	Code        string        `bson:"code" json:"code" valid:"required"`
	Description string        `bson:"description" json:"description" valid:"required"`
	Details     []*ItemDetail `bson:"details,omitempty" json:"details,omitempty" valid:"-"`
}

// AddDetail to item
func (c *Controller) AddDetail() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		itemID := mux.Vars(req)["id"]
		if len(itemID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		detail := &ItemDetail{}
		err := u.UnmarshalWithValidation(detail)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemsCollection)

		item := bson.M{"_id": itemID}
		record := bson.M{"$push": bson.M{"details": detail}}

		err = collection.Update(item, record)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("success", http.StatusCreated)
	})
}

// GetItemDetailsByID ...
func (c *Controller) GetItemDetailsByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		itemID := mux.Vars(req)["id"]

		if len(itemID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		var dto DetailResponseDto
		collection := session.DB(c.databaseName).C(ItemsCollection)
		err := collection.
			FindId(bson.ObjectIdHex(itemID)).
			Select(bson.M{"code": 1, "description": 1, "details": 1}).
			One(&dto)

		if err != nil {
			u.WriteJSONError("Something wrong, please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(dto)
	})
}
