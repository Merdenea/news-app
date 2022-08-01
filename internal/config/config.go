package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	configFileLocation = "./config.yaml"
)

type Config struct {
	NewsSources []string      `yaml:"news_sources"`
	Port        string        `yaml:"port"`
	CacheTTL    time.Duration `yaml:"cache_ttl"`
	DB          *Database     `yaml:"db"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func LoadConfig() (*Config, error) {
	configFile, err := ioutil.ReadFile(configFileLocation)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
