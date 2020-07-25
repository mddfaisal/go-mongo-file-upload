package router

import (
	"services/routes"

	"github.com/gorilla/mux"
)

// NewRouter new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.AllRoutes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}
	return router
}
