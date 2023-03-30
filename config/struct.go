package config

type Cfg struct {
	Server ServerCfg `yaml:"Server"`
	Mqtt   MqttCfg   `yaml:"Mqtt"`
	Jwt    JwtCfg    `yaml:"Jwt"`
	Door   DoorCfg   `yaml:"Door"`
}

type ServerCfg struct {
	Listen string `yaml:"Listen"`
	Port   int    `yaml:"Port"`
	Passwd string `yaml:"Password"`
}

type JwtCfg struct {
	Secret string `yaml:"Secret"`
}

type MqttCfg struct {
	Addr     string `yaml:"Addr"`
	Topic    string `yaml:"Topic"`
	ClientID string `yaml:"ClientID"`
}

type DoorCfg struct {
	OpenDelay int `yaml:"OpenDelay"`
	OpenAngle int `yaml:"OpenAngle"`
}
