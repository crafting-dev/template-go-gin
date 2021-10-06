package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error is a custom error structure.
type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// renderError renders a custom error, with set indentation.
func renderError(c *gin.Context, indent bool, err Error) {
	if indent {
		c.IndentedJSON(err.Code, gin.H{"error": err})
	} else {
		c.JSON(err.Code, gin.H{"error": err})
	}
}

// BadRequest responds with error status code 400, Bad Request.
func BadRequest(c *gin.Context) {
	var err Error
	err.Code = http.StatusBadRequest
	err.Status = "Bad Request"
	err.Message = "Bad request."
	renderError(c, false, err)
}

// NotFound responds with error status code 404, Not Found.
func NotFound(c *gin.Context) {
	var err Error
	err.Code = http.StatusNotFound
	err.Status = "Not Found"
	err.Message = "Not found."
	renderError(c, false, err)
}

// InternalServerError responds with error status code 500, Internal Server Error.
func InternalServerError(c *gin.Context) {
	var err Error
	err.Code = http.StatusInternalServerError
	err.Status = "Internal Server Error"
	err.Message = "Internal server error."
	renderError(c, false, err)
}
