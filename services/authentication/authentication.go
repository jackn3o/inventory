package authentication

import (
	"path"

	configuration "../../base/configuration"
	"../../base/connector"
	base "../../base/service"
	controller "../../controllers/authentication"
)

// authenticationService is one router service.
type authenticationService struct {
	store  *connector.Store
	config configuration.Config
}

func (s *authenticationService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.store, s.config)
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, ""),
			Handler: base.Use(c.GenerateToken(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, ""),
			Handler: base.Use(c.Authenticate(), mw...),
			Method:  "POST",
		},
	}
}

// New is a constructor for creating authentication service.
func New(store *connector.Store, config configuration.Config) base.Service {
	return &authenticationService{
		store:  store,
		config: config,
	}
}
