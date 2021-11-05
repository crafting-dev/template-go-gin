package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	Route{http.MethodGet, "/ping", Ping},
}
