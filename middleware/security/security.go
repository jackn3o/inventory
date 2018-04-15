package security

import (
	"net/http"

	base "../../base/service"
)

type middleware struct{}

func (mw *middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// New instantiate server side token checking middleware
func New() base.Middleware {
	return &middleware{}
}
