package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config is a struct for config
type Config struct {
	App AppConfig
	DB  DBConfig
}

// AppConfig is a struct for app config
type AppConfig struct {
	Host   string
	Port   string
	Secret string
}

// DBConfig is a struct for db config
type DBConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

var config Config

func init() {
	basePath, _ := os.Getwd()

	configPath := filepath.Join(basePath, "config", "config.json")

	byt, _ := ioutil.ReadFile(configPath)

	_ = json.Unmarshal(byt, &config)

	if config.App.Host == "" {
		config.App.Host = "0.0.0.0"
	}

	if config.App.Port == "" {
		config.App.Port = "8000"
	}
}

func GetConfig() Config {
	return config
}
