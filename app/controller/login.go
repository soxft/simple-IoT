package controller

import (
	"door/config"
	"door/lib/apiutil"
	"door/lib/jwtutil"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	api := apiutil.New(c)

	password := c.PostForm("password")
	if password == config.Server.Passwd {
		token, err := jwtutil.GenerateJwt(c.ClientIP())
		if err != nil {
			log.Println(err)
			api.Fail("login failed")
		}
		api.SuccessWithData("登陆成功", gin.H{
			"token": token,
		})
	} else {
		api.Fail("用户名或密码错误")
	}
}

func CheckPermission(c *gin.Context) {
	api := apiutil.New(c)

	api.Success("ok")
}
