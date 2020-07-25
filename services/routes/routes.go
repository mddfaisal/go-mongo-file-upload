package routes

import (
	"net/http"
	"services/mongocontroller"
)

// Route struct
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Routes struct
type Routes []Route

// AllRoutes all routes
var AllRoutes = Routes{
	Route{
		Name:    "getemail",
		Method:  "GET",
		Pattern: "/email",
		Handler: mongocontroller.GetEmail,
	},
	Route{
		Name:    "getemail",
		Method:  "POST",
		Pattern: "/email",
		Handler: mongocontroller.NewEmail,
	},
	Route{
		Name:    "getemail",
		Method:  "PUT",
		Pattern: "/email",
		Handler: mongocontroller.UpdateEmail,
	},
	Route{
		Name:    "getemail",
		Method:  "DELETE",
		Pattern: "/email",
		Handler: mongocontroller.DeleteEmail,
	},
}
