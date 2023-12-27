package config

type ConfigData struct {
	RunAddr  string
	BaseAddr string
}

func NewConfig() *ConfigData {
	return &ConfigData{}
}
