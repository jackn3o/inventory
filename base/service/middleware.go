package service

import "net/http"

// Middleware defines the callable that can be chained
type Middleware interface {
	// Handler will execute `next` appropriately after examining the current request
	// and often enhance the `context.Context` object passed to `http.Request`
	// to provide scoped values
	Handler(next http.Handler) http.Handler
}

// MiddlewareFunc provides the adapter for function middleware
type MiddlewareFunc func(http.Handler) http.Handler

// Handler function
func (mw MiddlewareFunc) Handler(next http.Handler) http.Handler {
	return mw(next)
}
