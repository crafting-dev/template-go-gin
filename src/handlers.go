package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Render applies indentation settings to response.
func Render(c *gin.Context, indent bool, obj interface{}) {
	if indent {
		c.IndentedJSON(http.StatusOK, obj)
	} else {
		c.JSON(http.StatusOK, obj)
	}
}

// Template Home route handler.
// Handles GET requests for root "/" path.
func TemplateHome(c *gin.Context) {
	Render(c, false, "Successfully launched!")
}
