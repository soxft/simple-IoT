package controller

import (
	"door/lib/apiutil"
	"door/process/mqtt"
	"github.com/gin-gonic/gin"
)

var lightStatus = false

func GetLight(c *gin.Context) {
	api := apiutil.New(c)
	if lightStatus {
		api.SuccessWithData("on", gin.H{
			"status": 1,
		})
	} else {
		api.SuccessWithData("off", gin.H{
			"status": 0,
		})
	}
}

func SetLight(c *gin.Context) {
	status := c.Param("status")
	if status == "1" {
		lightStatus = true
		mqtt.Publish("on")
	} else {
		lightStatus = false
		mqtt.Publish("off")
	}

	api := apiutil.New(c)
	api.Success("ok")
}
