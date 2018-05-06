package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
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

		isExist, err := c.isDescriptionExist(color.Description)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if isExist {
			u.WriteJSONError("Description exist", http.StatusConflict)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()
		collection := session.DB(c.databaseName).C(ColorSettingCollection)

		color.ID = bson.NewObjectId()
		timeNow := time.Now()
		color.CreatedDate = &timeNow
		color.CreatedBy = "todo"
		err = collection.Insert(color)

		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("New color successfully added: "+color.Description, http.StatusCreated)
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
			Select(bson.M{"description": 1, "hex": 1}).
			Sort("description").
			All(&colors)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(colors)
	})
}

// GetColorByID ...
func (c *Controller) GetColorByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		colorID := mux.Vars(req)["id"]

		if len(colorID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()
		collection := session.DB(c.databaseName).C(ColorSettingCollection)

		var dto Color
		err := collection.
			FindId(bson.ObjectIdHex(colorID)).
			Select(bson.M{"description": 1, "hex": 1}).
			One(&dto)

		if err != nil {
			u.WriteJSONError("Something wrong, please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(dto)
	})
}

// UpdateColor by id
func (c *Controller) UpdateColor() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		colorID := mux.Vars(req)["id"]

		if len(colorID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		color := &Color{}
		err := u.UnmarshalWithValidation(color)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		isRedefined, err := c.isDescriptionRedefine(colorID, color.Description)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if isRedefined {
			isExist, err := c.isDescriptionExist(color.Description)
			if err != nil {
				u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
				return
			}

			if isExist {
				u.WriteJSONError("description exist", http.StatusConflict)
				return
			}
		}

		session := c.store.DB.Copy()
		defer session.Close()
		collection := session.DB(c.databaseName).C(ColorSettingCollection)

		timeNow := time.Now()
		selector := bson.M{"_id": bson.ObjectIdHex(colorID)}
		updator := bson.M{"$set": bson.M{
			"description":  color.Description,
			"hex":          color.Hex,
			"modifiedDate": &timeNow,
			"modifiedBy":   "todo",
		}}
		if err := collection.Update(selector, updator); err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("Update Successful")
	})
}

// DeleteColor by id
func (c *Controller) DeleteColor() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		colorID := mux.Vars(req)["id"]

		if len(colorID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ColorSettingCollection)
		_, err := collection.RemoveAll(bson.M{"_id": bson.ObjectIdHex(colorID)})
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("Remove successful")
	})
}

func (c *Controller) isDescriptionRedefine(colorID string, requestedDescription string) (bool, error) {
	session := c.store.DB.Copy()
	defer session.Close()

	collection := session.DB(c.databaseName).C(ColorSettingCollection)
	var original Color
	err := collection.
		Find(bson.M{"_id": bson.ObjectIdHex(colorID)}).
		Select(bson.M{"description": 1}).
		One(&original)

	if err != nil {
		return false, err
	}

	if original.Description == requestedDescription {
		return false, nil
	}

	return true, err
}
func (c *Controller) isDescriptionExist(description string) (bool, error) {
	session := c.store.DB.Copy()
	defer session.Close()

	collection := session.DB(c.databaseName).C(ColorSettingCollection)
	count, err := collection.Find(bson.M{"description": description}).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
