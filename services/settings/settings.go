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
		base.EndPoint{
			Path:    path.Join(basePath, "/categories/{id}"),
			Handler: base.Use(c.GetCategoryByID(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/categories/{id}"),
			Handler: base.Use(c.UpdateCategory(), mw...),
			Method:  "PUT",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/categories/{id}"),
			Handler: base.Use(c.DeleteCategory(), mw...),
			Method:  "DELETE",
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

		// colors
		base.EndPoint{
			Path:    path.Join(basePath, "/colors"),
			Handler: base.Use(c.GetColors(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/colors"),
			Handler: base.Use(c.CreateColor(), mw...),
			Method:  "POST",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/colors/{id}"),
			Handler: base.Use(c.GetColorByID(), mw...),
			Method:  "GET",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/colors/{id}"),
			Handler: base.Use(c.UpdateColor(), mw...),
			Method:  "PUT",
		},
		base.EndPoint{
			Path:    path.Join(basePath, "/colors/{id}"),
			Handler: base.Use(c.DeleteColor(), mw...),
			Method:  "DELETE",
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
