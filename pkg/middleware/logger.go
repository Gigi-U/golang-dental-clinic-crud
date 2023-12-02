package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)
// Logger is a middleware function that logs information about incoming requests.
func Logger()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Capture the HTTP method, current time, requested URL, and initialize request size variable
		method:=ctx.Request.Method
		time:= time.Now()
		formattedTime := time.Format("02-01-2006 15:04:05")
		url:= ctx.Request.URL
		var size int
		// Proceed with the next middleware/handler in the chain
		ctx.Next()
		// If the response writer is not nil, capture the response size
		if ctx.Writer!= nil{
			size = ctx.Writer.Size()
		}
		// Print the request information to the console
		fmt.Printf("\nMethod:%s\nTime:%v\nPath:%s\nRequest size:%d\n", method, formattedTime, url, size)
	}
}