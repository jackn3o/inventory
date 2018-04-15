package authentication

import (
	"path"

	configuration "../../base/configuration"
	base "../../base/service"
	controller "../../controllers/authentication"
	"github.com/magicsoft-asia/vanda/lib/conn"
)

// authenticationService is one router service.
type authenticationService struct {
	conn   *conn.Bundle
	config configuration.Config
}

func (s *authenticationService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.config)
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, ""),
			Handler: base.Use(c.GenerateToken(), mw...),
			Method:  "GET",
		},
	}
}

// New is a constructor for creating authentication service.
func New(config configuration.Config) base.Service {
	return &authenticationService{
		config: config,
	}
}
