package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	JWT    JWT    `yaml:"jwt"`
}
