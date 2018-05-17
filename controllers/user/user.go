package user

import (
	"net/http"

	configuration "../../base/configuration"
	"../../base/connector"
	utility "../../base/utilities"
	"github.com/gorilla/mux"
)

// ResetPasswordDto ...
type ResetPasswordDto struct {
	Username        string `bson:"username" json:"username" valid:"required"`
	OldPassword     string `bson:"oldPassword" json:"oldPassword" valid:"required"`
	NewPassword     string `bson:"newPassword" json:"newPassword" valid:"required"`
	ConfirmPassword string `bson:"confirmPassword" json:"confirmPassword" valid:"required"`
}

// User dto
type User struct {
	Username    string `bson:"username" json:"username" valid:"required"`
	OldPassword string `bson:"password" json:"password" valid:"required"`
}

// Controller is return value for New method
type Controller struct {
	store  *connector.Store
	config configuration.Config
}

// New method is a constructor for controller.
func New(store *connector.Store, config configuration.Config) *Controller {
	return &Controller{
		store:  store,
		config: config,
	}
}

// ChangePassword login criteria
func (c *Controller) ChangePassword() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		username := mux.Vars(req)["username"]

		if len(username) <= 0 {
			u.WriteJSONError("username is require", http.StatusBadRequest)
			return
		}

		dto := &ResetPasswordDto{}
		err := u.UnmarshalWithValidation(dto)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		if dto.NewPassword != dto.ConfirmPassword {
			u.WriteJSONError("new password unmatch", http.StatusConflict)
			return
		}

		// userSession := c.store.DB.Copy()
		// defer userSession.Close()

		// // Generate "hash" from request password

		// masterDatabaseName := c.config.GetString(configuration.MasterDatabaseName)
		// userCollection := userSession.DB(masterDatabaseName).C(UsersCollection)
		// selector := bson.M{
		// 	"username": bson.RegEx{Pattern: dto.Username, Options: "i"},
		// }

		// var user User{}
		// err = userCollection.Find(selector).One(&user)
		// if err != nil {
		// 	u.WriteJSONError("User not found", http.StatusBadRequest)
		// 	return
		// }

		// isMatch := checkPasswordHash(dto.Password, user.Password)
		// if isMatch == false {
		// 	u.WriteJSONError("Password not match", http.StatusBadRequest)
		// 	return
		// }

		u.WriteJSON(dto)
	})
}

// constant key
const (
	UsersCollection = "users"
)
