package mqtt

import (
	"door/config"
	"testing"
)

func init() {
	config.CfgPath = "../../config.yaml"
	config.Init()
	return
}

func TestOn(t *testing.T) {
	Init()
	Publish("on")
	return
}

func TestOff(t *testing.T) {
	Init()
	Publish("off")
	return
}
