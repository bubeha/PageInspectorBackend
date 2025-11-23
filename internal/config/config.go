package config

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server      ServerConfig `json:"server"`
	Application AppConfig    `yaml:"application"`
}

type ServerConfig struct {
	Port    string `yaml:"port"`
	Timeout int16  `yaml:"timeout"`
}

type AppConfig struct {
	Environment string `yaml:"environment"`
}

func Load() *Config {
	data, err := os.ReadFile("./config/config.yml")

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)

	fmt.Printf("Port: %v\nTimeout: %v\n\n", cfg.Server.Port, cfg.Server.Timeout)

	return &cfg
}
