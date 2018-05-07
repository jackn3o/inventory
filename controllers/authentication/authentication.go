package authentication

// https://auth0.com/blog/authentication-in-golang/
// Import our dependencies. We'll use the standard http library as well as the gorilla router for this app
import (
	"fmt"
	"log"
	"net/http"
	"time"

	configuration "../../base/configuration"
	"../../base/connector"
	utility "../../base/utilities"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// LoginDto ...
type LoginDto struct {
	Username string `bson:"username" json:"username" valid:"require"`
	Password string `bson:"password" json:"password" valid:"require"`
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

// Authenticate login criteria
func (c *Controller) Authenticate() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		dto := &LoginDto{}
		err := u.UnmarshalWithValidation(dto)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		userSession := c.store.DB.Copy()
		defer userSession.Close()

		// Generate "hash" from request password
		hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			u.WriteJSONError("something wrong", http.StatusBadRequest)
		}

		userCollection := userSession.DB(c.config.GetString(configuration.MasterDatabaseName)).C("users")
		selector := bson.M{
			"username": bson.RegEx{Pattern: dto.Username, Options: "i"},
		}

		var user LoginDto
		err = userCollection.Find(selector).One(&user)
		if err != nil {
			u.WriteJSONError("User not found", http.StatusBadRequest)
			return
		}

		if err := bcrypt.CompareHashAndPassword(hash, []byte(user.Password)); err != nil {
			u.WriteJSONError("Invalid password", http.StatusBadRequest)
		}

		c.GenerateToken()

		u.WriteJSON("success")
	})
}

// GenerateToken for authentication
func (c *Controller) GenerateToken() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		/* Create the token */
		token := jwt.New(jwt.SigningMethodHS256)

		/* Create a map to store our claims */
		tokenClaims := token.Claims.(jwt.MapClaims)

		/* Set token claims */
		tokenClaims["admin"] = true
		tokenClaims["name"] = "admin"
		tokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate "hash" to store from user password
		hash, err := bcrypt.GenerateFromPassword([]byte("p@ssw0rd"), bcrypt.DefaultCost)
		if err != nil {
			// TODO: Properly handle error
			log.Fatal(err)
		}
		fmt.Println(string(hash))

		/* Sign the token with our secret */
		tokenString, _ := token.SignedString([]byte(c.config.GetString(configuration.SecretKey)))

		/* Finally, write the token to the browser window */
		u.WriteJSON(tokenString)
		// w.Write([]byte(tokenString))
	})
}
