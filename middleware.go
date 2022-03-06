package aproxiable

import (
	"github.com/gorilla/mux"
)

//standard middlewares

//authentication middleware

//AddMiddleware adds custom middleware
func addMiddlewares(middlewares ...mux.MiddlewareFunc) AproxiableOptionFunc {
	return func(a *Aproxiable) error {
		for _, middleware := range middlewares {
			a.middlewares = append(a.middlewares, middleware)
		}
		return nil
	}
}

//addReverseProxysToRouter adds a middleware to the mux router
func (a *Aproxiable) addMiddlewareToRouter(mwf ...mux.MiddlewareFunc) {
	a.router.Use(mwf...)
}
