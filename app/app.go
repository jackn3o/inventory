package app

import (
	"context"
	"fmt"
	"net/http"

	service "../base/service"
	authenticationService "../services/authentication"

	"../base/configuration"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// Start up for app
func Start(config configuration.Config) {
	rootContext, rootCancel := context.WithCancel(context.Background())

	// Here we are instantiating the gorilla/mux router

	appListenHost := config.GetString(configuration.AppListenHost)
	router := mux.NewRouter()

	setupRouter(rootContext, config, router)
	server := http.Server{
		ReadTimeout:  config.GetDuration(configuration.AppReadTimeout),
		WriteTimeout: config.GetDuration(configuration.AppWriteTimeout),
		Handler:      router,
		Addr:         appListenHost,
	}

	server.ListenAndServe()
	fmt.Printf("%s running on %s", "app", appListenHost)

	server.Shutdown(rootContext)
	rootCancel()
}

// setupRouter for api service
func setupRouter(ctx context.Context, config configuration.Config, router *mux.Router) {
	var jwtmw = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString(configuration.SecretKey)), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	apiRouter := router.PathPrefix(config.GetString(configuration.AppAPIBase)).Subrouter()

	service.Register(apiRouter, "/authentication", authenticationService.New(config), jwtmw)
	service.Register(apiRouter, "/public", authenticationService.New(config))
}
