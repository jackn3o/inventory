package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"gopkg.in/mgo.v2/bson"
)

// Outlet Model for settings.outlets Collection
type Outlet struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Code         string        `bson:"code" json:"code" valid:"required"`
	Description  string        `bson:"description" json:"description" valid:"required"`
	CreatedDate  *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// CreateOutlet in master data
func (c *Controller) CreateOutlet() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		outlet := &Outlet{}
		err := u.UnmarshalWithValidation(outlet)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}
		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(OutletSettingCollection)
		count, err := collection.Find(bson.M{"code": outlet.Code}).Count()
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if count != 0 {
			u.WriteJSONError("code exist", http.StatusConflict)
			return
		}

		outlet.ID = bson.NewObjectId()
		timeNow := time.Now()
		outlet.CreatedDate = &timeNow
		outlet.CreatedBy = "todo"
		err = collection.Insert(outlet)

		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			panic(err)
		}

		u.WriteJSON(outlet.ID, http.StatusCreated)
	})
}

// GetOutlets return list of outlets
func (c *Controller) GetOutlets() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		session := c.store.DB.Copy()
		defer session.Close()

		var outlets []Outlet
		collection := session.DB(c.databaseName).C(OutletSettingCollection)
		err := collection.
			Find(nil).
			Select(bson.M{"code": 1, "description": 1}).
			Sort("description").
			All(&outlets)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(outlets)
	})
}
