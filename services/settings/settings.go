package settings

import (
	"path"

	configuration "../../base/configuration"
	"../../base/connector"
	base "../../base/service"
	controller "../../controllers/settings"
)

// authenticationService is one router service.
type settingService struct {
	store  *connector.Store
	config configuration.Config
}

func (s *settingService) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	c := controller.New(s.store, s.config)
	return []base.EndPoint{
		// item
		base.EndPoint{
			Path:    path.Join(basePath, "/items"),
			Handler: base.Use(c.GetItems(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/items/{id}"),
			Handler: base.Use(c.GetItemByID(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/items"),
			Handler: base.Use(c.CreateItem(), mw...),
			Method:  "POST",
		},

		// category
		base.EndPoint{
			Path:    path.Join(basePath, "/categories"),
			Handler: base.Use(c.GetCategories(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/categories"),
			Handler: base.Use(c.CreateCategory(), mw...),
			Method:  "POST",
		},

		// outlet
		base.EndPoint{
			Path:    path.Join(basePath, "/outlets"),
			Handler: base.Use(c.GetOutlets(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/outlets"),
			Handler: base.Use(c.CreateOutlet(), mw...),
			Method:  "POST",
		},
	}
}

// New is a constructor for creating authentication service.
func New(store *connector.Store, config configuration.Config) base.Service {
	return &settingService{
		config: config,
		store:  store,
	}
}
