package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	CfgPath = "config.yaml"
	C       *Cfg
	Server  ServerCfg
	Mqtt    MqttCfg
	Jwt     JwtCfg
	Door    DoorCfg
)

func Init() {
	flag.StringVar(&CfgPath, "c", "config.yaml", "config file path")
	flag.Parse()

	C = &Cfg{}
	file, err := os.ReadFile(CfgPath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(file, C); err != nil {
		panic(err)
	}
	log.Println("Config loaded")
	log.Println(C)

	Server = C.Server
	Mqtt = C.Mqtt
	Jwt = C.Jwt
	Door = C.Door
}
