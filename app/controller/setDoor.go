package controller

import (
	"door/config"
	"door/lib/apiutil"
	"door/process/mqtt"
	"github.com/gin-gonic/gin"
	"time"
)

func SetDoor(c *gin.Context) {
	mqtt.Publish(config.Mqtt.Topic, "120")
	time.Sleep(time.Second * 2)
	mqtt.Publish(config.Mqtt.Topic, "0")

	api := apiutil.New(c)
	api.Success("ok")
}
