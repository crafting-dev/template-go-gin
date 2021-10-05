package api

import "github.com/gin-gonic/gin"

// HTTP methods.
const (
	GET    string = "GET"
	HEAD   string = "HEAD"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
	OPTION string = "OPTION"
	PATCH  string = "PATCH"
)

// A routing model.
type Route struct {
	Method   string
	Endpoint string
	Handler  gin.HandlerFunc
}

// A collection of routes.
type Routes []Route

var routes = Routes{
	Route{GET, "/", TemplateHome},
}
