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

// Handles GET requests for /ping?ping=<queryString>.
func Ping(c *gin.Context) {
	type Pong struct {
		Ping       string `json:"ping"`
		ReceivedAt string `json:"received_at"`
	}

	var pong Pong

	pong.Ping = c.Query("ping")
	if pong.Ping == "" {
		pong.Ping = "To ping, or not to ping; that is the question."
	}
	pong.ReceivedAt = time.Now().UTC().String()

	Render(c, false, pong)
}
