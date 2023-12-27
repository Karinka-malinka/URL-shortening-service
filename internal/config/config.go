package config

type ConfigData struct {
	RunAddr   string
	ShortAddr string
}

func NewConfig() *ConfigData {
	return &ConfigData{}
}
