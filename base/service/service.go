package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// EndPoint of specific operation
type EndPoint struct {
	Path    string
	Handler http.Handler
	Method  string
}

// Service is collection of endpoint
type Service interface {
	Create(path string, mw ...Middleware) []EndPoint
}

// Register service by grouped under same base path and middleware
func Register(router *mux.Router, basePath string, service Service, mw ...Middleware) {
	for _, ep := range service.Create(basePath, mw...) {
		router.Handle(ep.Path, ep.Handler).Methods(ep.Method)
	}
}

// Use the new handler by chaining middlewares against the handler
func Use(handler http.Handler, mw ...Middleware) http.Handler {

	h := handler

	// if chained with any middleware
	if len(mw) > 0 {

		n := len(mw) - 1

		// Call the right most middleware first
		for i := n; i >= 0; i-- {
			h = mw[i].Handler(h)
		}
	}

	return h
}
