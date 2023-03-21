package mqtt

import (
	"door/config"
	"testing"
	"time"
)

func init() {
	config.CfgPath = "../../config.yaml"
	config.Init()
	Init()
	return
}

func TestOn(t *testing.T) {
	Publish("on")
	time.Sleep(time.Second * 2)
	return
}

func TestOff(t *testing.T) {
	Publish("off")
	time.Sleep(time.Second * 2)
	return
}

//func TestGet(t *testing.T) {
//	Client.Publish(config.Mqtt.Topic, 0, false, "0")
//	return
//}
