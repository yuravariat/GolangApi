package main

import (
	"GoApi/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"PushRecieve",
		"POST",
		"/reciever",
		controllers.PushRecieve,
	},
	Route{
		"HotelsIndex",
		"GET",
		"/hotels",
		controllers.HotelsIndex,
	},
	Route{
		"HotelsFind",
		"GET",
		"/hotels/{id}",
		controllers.HotelsFind,
	},
	Route{
		"HotelsAdd",
		"POST",
		"/hotels",
		controllers.HotelsAdd,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controllers.TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controllers.TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controllers.TodoShow,
	},
}
