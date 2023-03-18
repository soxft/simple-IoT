package main

import (
	"door/config"
	"door/process/mqtt"
	"door/process/web"
)

func main() {
	config.Init()
	mqtt.Init()
	web.Init()
}
