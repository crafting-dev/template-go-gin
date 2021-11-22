package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/location"
)

// Router mode and server port info.
type Context struct {
	Mode string
	Port string
}

// Initialize server.
func Init(ctx Context) {
	// Set db credentials
	SetupMySQLConfig()

	// Set gin mode
	gin.SetMode(ctx.Mode)

	// Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	// configure to automatically detect scheme and host
  // - use http when default scheme cannot be determined
  // - use localhost:PORT when default host cannot be determined
  router.Use(location.New(location.Config{
		Scheme: "http",
		Host: "localhost:" + ctx.Port,
		Base: "/",
	}))

	// Handle api routes
	base := router.Group("/")
	for _, route := range routes {
		base.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match request url,
	// return 404 (Not Found)
	router.NoRoute(NotFound)

	// Listen and serve
	router.Run()
}
