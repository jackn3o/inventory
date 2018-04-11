package authentication

import (
	"net/http"
	"path"

	base "../../base/service"
)

type service struct {
}

func (s *service) CreateEndPoints(basePath string, mw ...base.Middleware) []base.EndPoint {
	return []base.EndPoint{
		base.EndPoint{
			Path:    path.Join(basePath, ""),
			Handler: base.Use(NotImplemented, mw...),
			Method:  "GET",
		},
	}
}

// func New() router.Service {
// 	return service{}
// }

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
