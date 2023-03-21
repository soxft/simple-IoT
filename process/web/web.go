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

	{
		engine.GET("/light", controller.GetLight)
		engine.POST("/light/:status", controller.SetLight)
	}

	{
		engine.POST("/door", controller.SetDoor)
	}

	{
		engine.GET("/device/last_seen", controller.GetDeviceLastOnline)
		engine.POST("/device/ping", controller.DoPing)
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
