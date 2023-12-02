package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
// response is a struct representing a generic JSON response format.
type response struct {
	Data interface{} `json:"data"`
	Message string   `json:"message"`
}
// errorResponse is a struct representing an error JSON response format.
type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
// Response sends a generic JSON response with the specified status code and data.
func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
// Success sends a success JSON response with the specified status code, data, and message.
func Success(c *gin.Context, status int, data interface{}, message string) {
	Response(c, status, response{Data: data, Message: message})
}

// Error sends an error JSON response with the specified status code, formatted message, and optional arguments.
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}