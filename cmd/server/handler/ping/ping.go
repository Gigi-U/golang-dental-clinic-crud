package ping

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Ping´s Controller Struct
type Controller struct {
}

func NewControllerPing() *Controller {
	return &Controller{}
}

// Method HandlerPing is the Ping´s handler
func (c *Controller) HandlerPing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	}
}