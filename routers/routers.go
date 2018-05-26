package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	handler "github.com/roshniashok/crud/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// abstracting route details as struct
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

//defining the route and handler
var routes = Routes{
	Route{
		"getRecord",
		"GET",
		"/record",
		handler.GetRecord,
	},
	Route{
		"postRecord",
		"POST",
		"/record",
		handler.PostRecord,
	},
}
