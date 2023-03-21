package controller

import (
	"door/config"
	"door/lib/apiutil"
	"door/process/mqtt"
	"github.com/gin-gonic/gin"
	"time"
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
		mqtt.Publish(config.Mqtt.Topic, "80")
		time.Sleep(time.Second)
		mqtt.Publish(config.Mqtt.Topic, "90")
	} else {
		lightStatus = false
		mqtt.Publish(config.Mqtt.Topic, "100")
		time.Sleep(time.Second)
		mqtt.Publish(config.Mqtt.Topic, "90")
	}

	api := apiutil.New(c)
	api.Success("ok")
}
