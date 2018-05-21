package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
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
			c.logger.Warn(err)
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}
		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(OutletSettingCollection)
		count, err := collection.Find(bson.M{"code": outlet.Code}).Count()
		if err != nil {
			c.logger.Error(err)
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
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
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
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(outlets)
	})
}

// GetOutletByID ...
func (c *Controller) GetOutletByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		outletID := mux.Vars(req)["id"]

		session := c.store.DB.Copy()
		defer session.Close()

		var outlet Outlet
		collection := session.DB(c.databaseName).C(OutletSettingCollection)
		err := collection.
			FindId(bson.ObjectIdHex(outletID)).
			Select(bson.M{"code": 1, "description": 1}).
			One(&outlet)
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(outlet)
	})
}

// UpdateOutlet by id
func (c *Controller) UpdateOutlet() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		outletID := mux.Vars(req)["id"]

		if len(outletID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		outlet := &Outlet{}
		err := u.UnmarshalWithValidation(outlet)
		if err != nil {
			c.logger.Warn(err)
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		isExist, err := c.isKeyFieldsExist(outletID, OutletSettingCollection, outlet.Code, outlet.Description)
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if isExist {
			u.WriteJSONError("Code or Description exist", http.StatusConflict)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()
		collection := session.DB(c.databaseName).C(OutletSettingCollection)

		timeNow := time.Now()
		selector := bson.M{"_id": bson.ObjectIdHex(outletID)}
		updator := bson.M{"$set": bson.M{
			"code":         outlet.Code,
			"description":  outlet.Description,
			"modifiedDate": &timeNow,
			"modifiedBy":   "todo",
		}}
		if err := collection.Update(selector, updator); err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("Update Successful")
	})
}

// DeleteOutlet by id
func (c *Controller) DeleteOutlet() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		outletID := mux.Vars(req)["id"]

		if len(outletID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(OutletSettingCollection)
		_, err := collection.RemoveAll(bson.M{"_id": bson.ObjectIdHex(outletID)})
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("Remove successful")
	})
}
