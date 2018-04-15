package authentication

// https://auth0.com/blog/authentication-in-golang/
// Import our dependencies. We'll use the standard http library as well as the gorilla router for this app
import (
	"net/http"
	"time"

	configuration "../../base/configuration"
	utility "../../base/utilities"
	jwt "github.com/dgrijalva/jwt-go"
)

// Controller is return value for New method
type Controller struct {
	// conn   *conn.Bundle
	config configuration.Config
}

// New method is a constructor for controller.
func New(config configuration.Config) *Controller {
	// conn *conn.Bundle,// todo add back to para after db is setup
	return &Controller{
		// conn: conn,
		config: config,
	}
}

// GenerateToken for authentication
func (c *Controller) GenerateToken() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)
		/* Create the token */
		token := jwt.New(jwt.SigningMethodHS256)

		/* Create a map to store our claims */
		claims := token.Claims.(jwt.MapClaims)

		/* Set token claims */
		claims["admin"] = true
		claims["name"] = "Ado Kukic"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		/* Sign the token with our secret */
		tokenString, _ := token.SignedString([]byte(c.config.GetString(configuration.SecretKey)))

		/* Finally, write the token to the browser window */
		u.WriteJSON(tokenString)
		// w.Write([]byte(tokenString))
	})
}
