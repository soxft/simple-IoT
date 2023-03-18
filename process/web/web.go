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

	engine.GET("/light", controller.GetLight)
	engine.POST("/light/:status", controller.SetLight)

	listenAddr := config.Server.Listen + ":" + strconv.Itoa(config.Server.Port)
	log.Println("Web server is listening on " + listenAddr)
	engine.Run(listenAddr)
}
