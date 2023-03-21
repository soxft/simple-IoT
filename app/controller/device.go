package controller

import (
	"door/config"
	"door/global"
	"door/lib/apiutil"
	"door/process/mqtt"
	"github.com/gin-gonic/gin"
	"time"
)

func DoPing(c *gin.Context) {
	mqtt.Publish(config.Mqtt.Topic, "ping")
	api := apiutil.New(c)
	api.Success("send ping to device done")
}

func GetDeviceLastOnline(c *gin.Context) {
	api := apiutil.New(c)
	api.SuccessWithData("ok", gin.H{
		"last_pong":     global.DeviceLastOnlineTime,
		"last_pong_str": time.Unix(global.DeviceLastOnlineTime, 0).Format("2006-01-02 15:04:05"),
	})
}
