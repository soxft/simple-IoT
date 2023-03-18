package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Server", "light-by-xcsoft")
		if c.Request.Method == "OPTIONS" {
			//	c.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, OPTIONS")
			//	c.Header("Access-Control-Max-Age", "86400")
			//	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.AbortWithStatus(204)
			return
		}
	}
}
