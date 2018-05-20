package user

import (
	"path"

	configuration "../../base/configuration"
	"../../base/connector"
	base "../../base/service"
	controller "../../controllers/authentication"
)

// userService is one router service.
type userService struct {
	store  *connector.Store
	config configuration.Config
}

func (s *userService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.store, s.config)
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, "/{username}/password"),
			Handler: base.Use(c.ChangePassword(), mw...),
			Method:  "PUT",
		},
	}
}

// New is a constructor for creating authentication service.
func New(store *connector.Store, config configuration.Config) base.Service {
	return &userService{
		store:  store,
		config: config,
	}
}
