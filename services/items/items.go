package items

import (
	"path"

	configuration "../../base/configuration"
	"../../base/connector"
	base "../../base/service"
	controller "../../controllers/items"
)

// itemService is one router service.
type itemService struct {
	store  *connector.Store
	config configuration.Config
}

func (s *itemService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.store, s.config)
	return []base.EndPoint{
		// items
		base.EndPoint{
			Path:    path.Join(basePath, "/"),
			Handler: base.Use(c.GetItems(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/"),
			Handler: base.Use(c.CreateItem(), mw...),
			Method:  "POST",
		},

		// details
		base.EndPoint{
			Path:    path.Join(basePath, "/{id}"),
			Handler: base.Use(c.GetItemDetailsByID(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/{id}/details"),
			Handler: base.Use(c.AddDetail(), mw...),
			Method:  "POST",
		},
	}
}

// New is a constructor for creating authentication service.
func New(store *connector.Store, config configuration.Config) base.Service {
	return &itemService{
		config: config,
		store:  store,
	}
}
