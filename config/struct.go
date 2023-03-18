package config

type Cfg struct {
	Server ServerCfg `yaml:"Server"`
	Mqtt   MqttCfg   `yaml:"Mqtt"`
}

type ServerCfg struct {
	Listen string `yaml:"Listen"`
	Port   int    `yaml:"Port"`
}

type MqttCfg struct {
	Addr     string `yaml:"Addr"`
	Topic    string `yaml:"Topic"`
	ClientID string `yaml:"ClientID"`
}
