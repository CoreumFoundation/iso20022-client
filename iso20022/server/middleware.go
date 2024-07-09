package server

import (
	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

// CORSMiddleware provides CORS ability for the router
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, HEAD")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// InjectDependencies injects message parser and channels
func InjectDependencies(parser processes.Parser, messageQueue processes.MessageQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("parser", parser)
		c.Set("messageQueue", messageQueue)

		c.Next()
	}
}
