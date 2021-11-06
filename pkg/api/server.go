package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// Handle api routes
	base := router.Group("/")
	for _, route := range routes {
		base.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match request url,
	// return 404 (Not Found)
	router.NoRoute(NotFound)

	// Listen and serve on 0.0.0.0:PORT
	router.Run(":" + ctx.Port)
}
