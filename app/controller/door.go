package controller

import (
	"door/config"
	"door/lib/apiutil"
	"door/process/mqtt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func SetDoor(c *gin.Context) {
	api := apiutil.New(c)
	api.Success("ok")

	go func() {
		mqtt.Publish(config.Mqtt.Topic, strconv.Itoa(config.Door.OpenAngle))
		time.Sleep(time.Millisecond * time.Duration(config.Door.OpenDelay))
		mqtt.Publish(config.Mqtt.Topic, "0")
	}()
}
