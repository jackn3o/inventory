package settings

import (
	configuration "../../base/configuration"
	"../../base/connector"
)

// Controller is return value for New method
type Controller struct {
	databaseName string
	store        *connector.Store
	config       configuration.Config
}

// New method is a constructor for controller.
func New(store *connector.Store, config configuration.Config) *Controller {
	return &Controller{
		databaseName: config.GetString(configuration.DatabaseName),
		store:        store,
		config:       config,
	}
}

// Key for collection
const (
	ItemSettingCollection     = "settings.items" // main setting
	CategorySettingCollection = "settings.categories"
	ColorSettingCollection    = "settings.colors"
	OutletSettingCollection   = "settings.outlets"
)
