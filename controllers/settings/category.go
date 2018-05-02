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

		u.WriteJSON(category.ID, http.StatusCreated)
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
