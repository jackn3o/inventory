package settings

import (
	configuration "../../base/configuration"
	"../../base/connector"
	"../../base/log"
)

// Controller is return value for New method
type Controller struct {
	databaseName string
	store        *connector.Store
	config       configuration.Config
	logger       log.Logger
}

// New method is a constructor for controller.
func New(store *connector.Store, config configuration.Config) *Controller {
	return &Controller{
		databaseName: config.GetString(configuration.MainDatabaseName),
		store:        store,
		config:       config,
		logger:       log.New(),
	}
}

// Key for collection
const (
	CategorySettingCollection = "settings.categories"
	ColorSettingCollection    = "settings.colors"
	OutletSettingCollection   = "settings.outlets"
)
