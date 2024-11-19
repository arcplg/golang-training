package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Firebase struct {
		ServiceAccountKeyPath string `yaml:"serviceAccountKeyPath"`
	} `yaml:"firebase"`
	App struct {
		Port  int  `yaml:"port"`
		Debug bool `yaml:"debug"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	configPath := "./config.yaml"
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
