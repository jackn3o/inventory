package authentication

// https://auth0.com/blog/authentication-in-golang/
// Import our dependencies. We'll use the standard http library as well as the gorilla router for this app
import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	"../../base/connector"
	"../../base/log"
	utility "../../base/utilities"
	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// constant key
const (
	UsersCollection = "users"
)

// LoginDto ...
type LoginDto struct {
	Username string `bson:"username" json:"username" valid:"required"`
	Password string `bson:"password" json:"password" valid:"required"`
}

// User model
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	Username     string        `bson:"username" json:"username" valid:"required"`
	Password     string        `bson:"password" json:"password" valid:"required"`
	CreatedDate  *time.Time    `bson:"createdDate"`
	CreatedBy    string        `bson:"createBy"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty"`
}

// TokenResponseDto ...
type TokenResponseDto struct {
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}

// Controller is return value for New method
type Controller struct {
	store  *connector.Store
	config configuration.Config
	logger log.Logger
}

// New method is a constructor for controller.
func New(store *connector.Store, config configuration.Config) *Controller {
	return &Controller{
		store:  store,
		config: config,
		logger: log.New("authentication-controller"),
	}
}

// Authenticate login criteria
func (c *Controller) Authenticate() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		dto := &LoginDto{}
		err := u.UnmarshalWithValidation(dto)
		if err != nil {
			c.logger.Warn(err)
			u.WriteJSONError(err, http.StatusBadRequest)
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

		var user LoginDto
		err = userCollection.Find(selector).One(&user)
		if err != nil {
			if dto.Username == "admin" {
				user, err := c.createDefaultAdminUser(masterDatabaseName)
				if err != nil {
					c.logger.Warn(err)
					u.WriteJSONError(err, http.StatusBadRequest)
					return
				}
				//todo use user to login

			} else {
				c.logger.Error(err)
				u.WriteJSONError("User not found", http.StatusBadRequest)
				return
			}
		}

		isMatch := checkPasswordHash(dto.Password, user.Password)
		if isMatch == false {
			u.WriteJSONError("Password not match", http.StatusBadRequest)
			return
		}

		expiredAt := time.Now().Add(time.Hour * 24).Unix()

		responseDTO := TokenResponseDto{
			Username:  dto.Username,
			Token:     c.generateToken(dto.Username, expiredAt),
			ExpiredAt: expiredAt,
		}

		u.WriteJSON(responseDTO)
	})
}

// generateToken for authentication
func (c *Controller) generateToken(username string, expiredAt int64) string {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	tokenClaims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	tokenClaims["admin"] = true
	tokenClaims["name"] = username
	tokenClaims["exp"] = expiredAt

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString([]byte(c.config.GetString(configuration.SecretKey)))

	return tokenString
}

func (c *Controller) createDefaultAdminUser(masterDatabaseName string) (User, error) {
	var user User

	defaultPassword, err := hashPassword("p@ssw0rd")
	if err != nil {
		return user, err
	}

	timeNow := time.Now()
	user = User{
		ID:          bson.NewObjectId(),
		Username:    "admin",
		Password:    defaultPassword,
		CreatedBy:   "system",
		CreatedDate: &timeNow,
	}

	userSession := c.store.DB.Copy()
	userCollection := userSession.DB(masterDatabaseName).C(UsersCollection)

	err = userCollection.Insert(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
