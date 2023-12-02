package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
// Authenticate is a middleware function that checks the presence and validity of a token in the request header.
func Authenticate()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Retrieve the token from the request header and the expected token from the environment
		tokenHeader := ctx.GetHeader("TokenPostman")
		tokenEnv:=os.Getenv("TOKEN")
		// Check if the token is missing or invalid
		if(tokenHeader == "" || tokenHeader!= tokenEnv){
			// Abort the request with a 401 Unauthorized status and a JSON response
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}else{
			// If the token is valid, proceed to the next middleware/handler in the chain
			ctx.Next()
		}
	}
}