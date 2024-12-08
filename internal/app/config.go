package app

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host             string   `yaml:"host" json:"host" env:"SERVER_HOST"`
	Port             string   `yaml:"port" json:"port" env:"SERVER_PORT"`
	UserHost         string   `yaml:"user_host" json:"user_host" env:"USER_HOST"`
	UserInternalPort string   `yaml:"user_internal_port" json:"user_internal_port" env:"USER_INTERNAL_PORT"`
	DB               Database `yaml:"database" json:"database"`
}

type Database struct {
	DSN string `yaml:"dsn" json:"dsn"`
}

func NewConfig(configPath string) (*Config, error) {
	var config = new(Config)

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
