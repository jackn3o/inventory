package authentication

import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	utility "../../base/utilities"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// ResetPasswordDto ...
type ResetPasswordDto struct {
	Username        string `bson:"username" json:"username" valid:"-"`
	OldPassword     string `bson:"oldPassword" json:"oldPassword" valid:"required"`
	NewPassword     string `bson:"newPassword" json:"newPassword" valid:"required"`
	ConfirmPassword string `bson:"confirmPassword" json:"confirmPassword" valid:"required"`
}

// User dto
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Username string        `bson:"username" json:"username" valid:"required"`
	Password string        `bson:"password" json:"password" valid:"required"`
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
			c.logger.Warn(err)
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		if dto.NewPassword != dto.ConfirmPassword {
			u.WriteJSONError("new password unmatch", http.StatusConflict)
			return
		}

		userSession := c.store.DB.Copy()
		defer userSession.Close()

		// Generate "hash" from request password

		masterDatabaseName := c.config.GetString(configuration.MasterDatabaseName)
		userCollection := userSession.DB(masterDatabaseName).C(UsersCollection)
		selector := bson.M{
			"username": bson.RegEx{Pattern: dto.Username, Options: "i"},
		}

		var user User
		err = userCollection.Find(selector).One(&user)
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("User not found", http.StatusBadRequest)
			return
		}

		isMatch := checkPasswordHash(dto.OldPassword, user.Password)
		if isMatch == false {
			u.WriteJSONError("Old Password invalid", http.StatusBadRequest)
			return
		}

		newPassword, err := hashPassword(dto.NewPassword)
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusBadRequest)
			return
		}

		timeNow := time.Now()
		updator := bson.M{"$set": bson.M{
			"password":     newPassword,
			"modifiedDate": &timeNow,
		}}

		if err := userCollection.Update(selector, updator); err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON("success")
	})
}
