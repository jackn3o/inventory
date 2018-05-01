package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"gopkg.in/mgo.v2/bson"
)

// Color Model for settings.colors Collection
type Color struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Description  string        `bson:"description" json:"description" valid:"required"`
	Hex          string        `bson:"hex" json:"hex" valid:"required"`
	CreatedDate  *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// CreateColor in master data
func (c *Controller) CreateColor() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		color := &Color{}
		err := u.UnmarshalWithValidation(color)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}
		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ColorSettingCollection)
		count, err := collection.Find(bson.M{"code": color.Description}).Count()
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if count != 0 {
			u.WriteJSONError("code exist", http.StatusConflict)
			return
		}

		color.ID = bson.NewObjectId()
		timeNow := time.Now()
		color.CreatedDate = &timeNow
		color.CreatedBy = "todo"
		err = collection.Insert(color)

		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			panic(err)
		}

		u.WriteJSON(color.ID, http.StatusCreated)
	})
}

// GetColors return list of outlets
func (c *Controller) GetColors() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		session := c.store.DB.Copy()
		defer session.Close()

		var colors []Color
		collection := session.DB(c.databaseName).C(ColorSettingCollection)
		err := collection.
			Find(nil).
			Select(bson.M{"code": 1, "description": 1}).
			Sort("description").
			All(&colors)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(colors)
	})
}
