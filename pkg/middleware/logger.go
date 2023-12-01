package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		method:=ctx.Request.Method
		time:= time.Now()
		formattedTime := time.Format("02-01-2006 15:04:05")
		url:= ctx.Request.URL
		var size int

		ctx.Next()

		if ctx.Writer!= nil{
			size = ctx.Writer.Size()
		}

		fmt.Printf("\nMethod:%s\nTime:%v\nPath:%s\nRequest size:%d\n", method, formattedTime, url, size)
	}
}