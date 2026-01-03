package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		end := time.Since(start)

		log.Printf("[GIN-CUSTOM] %s | %d | %s | %s | %s %s\n",
            end.String(),
            ctx.Writer.Status(),
            ctx.ClientIP(),
            ctx.Request.Method,
            ctx.Request.URL.Path,
            ctx.Request.Proto,
        )
	}
}