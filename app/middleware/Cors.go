package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Server", "light-by-xcsoft")
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "OPTIONS,"+c.Request.Header.Get("Access-Control-Request-Method"))
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Access-Control-Allow-Headers", "DNT,Content-Type,authorization")
			c.AbortWithStatus(204)
			return
		}
	}
}
