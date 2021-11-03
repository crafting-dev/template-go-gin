package api

import (
	"net/http"
	"time"

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

// Handles GET requests for /ping path.
// Accepts a 'ping' query string parameter.
// If no query string param is provided,
// it responds with a default string.
// Otherwise, it responds with the query
// ping param and current timestamp.
func PingPong(c *gin.Context) {
	type Response struct {
		Ping       string `json:"ping"`
		ReceivedAt string `json:"received_at"`
	}

	var response Response

	response.Ping = c.DefaultQuery("ping", "To ping, or not to ping; that is the question.")
	response.ReceivedAt = time.Now().UTC().String()

	Render(c, false, response)
}
