package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	service "../base/service"
	authenticationservice "../services/authentication"
	itemsservice "../services/items"
	settingsservice "../services/settings"
	userservice "../services/user"

	"../base/configuration"
	"../base/connector"
	"../base/log"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Start up for app
func Start(config configuration.Config) {
	rootContext, rootCancel := context.WithCancel(context.Background())
	logger := log.New("app")

	appListenHost := config.GetString(configuration.AppListenHost)
	router := mux.NewRouter()

	logger.Info("action=setup-router")
	setupRouter(rootContext, config, router)

	logger.Info("action=setup-cors")
	corsmw := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Origin", "Authorization", "X-Client"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	server := http.Server{
		ReadTimeout:  config.GetDuration(configuration.AppReadTimeout) * time.Second,
		WriteTimeout: config.GetDuration(configuration.AppWriteTimeout) * time.Second,
		Handler:      corsmw.Handler(router),
		Addr:         appListenHost,
	}

	go func() {
		logger.Fatal(server.ListenAndServe())
	}()

	// setup logger
	// verify log level value
	configLogLevel := config.Get(configuration.LogLevel)

	switch configLogLevel.(type) {
	case log.LogLevel:
		log.SetLogLevel(configLogLevel.(log.LogLevel))
		break
	case int:
		log.SetLogLevel(log.LogLevel(configLogLevel.(int)))
		break
	}

	for {
		logger.Infof("%s running on %s", "app", appListenHost)
		sig := waitForSignals()

		// must survive sighub signal
		if sig == syscall.SIGHUP {
			continue
		}

		break
	}

	logger.Info("shutting down")

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

	store := connector.New(config)

	apiRouter := router.PathPrefix(config.GetString(configuration.AppAPIBase)).Subrouter()
	service.Register(apiRouter, "/authentication", authenticationservice.New(store, config))
	service.Register(apiRouter, "/users", userservice.New(store, config), jwtmw)
	service.Register(apiRouter, "/items", itemsservice.New(store, config), jwtmw)
	service.Register(apiRouter, "/settings", settingsservice.New(store, config), jwtmw)

}

func waitForSignals() os.Signal {
	sink := make(chan os.Signal, 1)
	defer close(sink)

	// wait for signal
	signal.Notify(sink, signals...)
	defer signal.Ignore(signals...)

	return <-sink
}
