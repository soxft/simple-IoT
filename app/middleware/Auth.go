package middleware

import (
	"door/lib/apiutil"
	"door/lib/jwtutil"
	"github.com/gin-gonic/gin"
)

func AuthPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := apiutil.New(c)

		// check jwt token
		var token string
		if token = jwtutil.GetJwtFromAuth(c.GetHeader("Authorization")); token == "" {
			api.Abort401("Unauthorized", 0)
			return
		}
		if err := jwtutil.CheckPermission(token); err != nil {
			api.Abort401("Unauthorized", 1)
			return
		}

		c.Next()
	}
}
