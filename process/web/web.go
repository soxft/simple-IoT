package web

import (
	"door/app/controller"
	"door/app/middleware"
	"door/config"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Init() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(middleware.Cors())

	light := engine.Group("/light")
	light.Use(middleware.AuthPermission())
	{
		light.GET("", controller.GetLight)
		light.POST("/:status", controller.SetLight)
	}

	door := engine.Group("/door")
	door.Use(middleware.AuthPermission())
	{
		door.POST("", controller.SetDoor)
	}

	device := engine.Group("/device")
	device.Use(middleware.AuthPermission())
	{
		device.GET("/last_seen", controller.GetDeviceLastOnline)
		device.POST("/ping", controller.DoPing)
	}

	auth := engine.Group("/auth")
	{
		engine.POST("/login", controller.Login)
		auth.Use(middleware.AuthPermission())
		auth.GET("/permission", controller.CheckPermission)
	}

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/web")
	})

	engine.Static("/_next", "public/_next")
	engine.Static("/web", "public")

	listenAddr := config.Server.Listen + ":" + strconv.Itoa(config.Server.Port)
	log.Println("Web server is listening on " + listenAddr)
	engine.Run(listenAddr)
}
