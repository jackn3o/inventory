package settings

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Category Model for settings.categories Collection
type Category struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty" valid:"-"`
	Code         string        `bson:"code" json:"code" valid:"required"`
	Description  string        `bson:"description" json:"description" valid:"required"`
	CreatedDate  *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// CreateCategory in master data
func (c *Controller) CreateCategory() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		category := &Category{}
		err := u.UnmarshalWithValidation(category)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		isExist, err := c.isKeyFieldsExist("", CategorySettingCollection, category.Code, category.Description)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if isExist {
			u.WriteJSONError("Code or Description exist", http.StatusConflict)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(CategorySettingCollection)
		count, err := collection.Find(bson.M{"code": category.Code}).Count()
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if count != 0 {
			u.WriteJSONError("code exist", http.StatusConflict)
			return
		}

		category.ID = bson.NewObjectId()
		timeNow := time.Now()
		category.CreatedDate = &timeNow
		category.CreatedBy = "todo"
		err = collection.Insert(category)

		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("New category successfully added: "+category.Description, http.StatusCreated)
	})
}

// GetCategories return list of category
func (c *Controller) GetCategories() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		vars := mux.Vars(req)
		sortParam := "description"
		sortAscending, ok := vars["ascending"]
		if ok {
			if sortAscending == "false" {
				sortParam = "-description"
			}
		}

		session := c.store.DB.Copy()
		defer session.Close()

		var categories []Category
		collection := session.DB(c.databaseName).C(CategorySettingCollection)
		err := collection.
			Find(nil).
			Select(bson.M{"code": 1, "description": 1}).
			Sort(sortParam).
			All(&categories)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(categories)
	})
}

// GetCategoryByID ...
func (c *Controller) GetCategoryByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		categoryID := mux.Vars(req)["id"]

		session := c.store.DB.Copy()
		defer session.Close()

		var categories Category
		collection := session.DB(c.databaseName).C(CategorySettingCollection)
		err := collection.
			FindId(bson.ObjectIdHex(categoryID)).
			Select(bson.M{"code": 1, "description": 1}).
			One(&categories)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(categories)
	})
}

// UpdateCategory by id
func (c *Controller) UpdateCategory() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		categoryID := mux.Vars(req)["id"]

		if len(categoryID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		category := &Category{}
		err := u.UnmarshalWithValidation(category)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		isExist, err := c.isKeyFieldsExist(categoryID, CategorySettingCollection, category.Code, category.Description)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		if isExist {
			u.WriteJSONError("Code or Description exist", http.StatusConflict)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()
		collection := session.DB(c.databaseName).C(CategorySettingCollection)

		timeNow := time.Now()
		selector := bson.M{"_id": bson.ObjectIdHex(categoryID)}
		updator := bson.M{"$set": bson.M{
			"code":         category.Code,
			"description":  category.Description,
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

// DeleteCategory by id
func (c *Controller) DeleteCategory() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		categoryID := mux.Vars(req)["id"]

		if len(categoryID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(CategorySettingCollection)
		_, err := collection.RemoveAll(bson.M{"_id": bson.ObjectIdHex(categoryID)})
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("Remove successful")
	})
}
