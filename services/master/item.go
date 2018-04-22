package master

import (
	"path"

	configuration "../../base/configuration"
	"../../base/connector"
	base "../../base/service"
	controller "../../controllers/master"
)

// authenticationService is one router service.
type masterService struct {
	store  *connector.Store
	config configuration.Config
}

func (s *masterService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.store, s.config)
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, "/items"),
			Handler: base.Use(c.CreateItem(), mw...),
			Method:  "POST",
		},
	}
}

// New is a constructor for creating authentication service.
func New(store *connector.Store, config configuration.Config) base.Service {
	return &masterService{
		config: config,
		store:  store,
	}
}
