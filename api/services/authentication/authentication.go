package authentication

import (
	"net/http"
	"path"

	base "../../base/service"
	utility "../../base/utilities"
)

type service struct {
}

type test struct {
}

func (s *service) Create(basePath string, mw ...base.Middleware) []base.EndPoint {
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, ""),
			Handler: base.Use(NotImplemented, mw...),
			Method:  "GET",
		},
	}
}

// New is a constructor for creating authentication service.
func New() base.Service {
	return &service{}
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	u := utility.New(w, r)

	u.WriteJSON("Not Implemented")
})
