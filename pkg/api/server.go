package api

import "github.com/gin-gonic/gin"

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

	// Handle versioned api routes
	v1 := router.Group("/v1")
	for _, route := range routes {
		v1.Handle(route.Method, route.Endpoint, route.Handler)
	}

	// If no routers match request url,
	// return 404 (Not Found)
	router.NoRoute(NotFound)

	// Listen and serve
	router.Run(":" + ctx.Port)
}
