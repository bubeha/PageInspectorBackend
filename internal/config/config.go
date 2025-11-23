package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server      ServerConfig `json:"server"`
	Application AppConfig    `yaml:"application"`
	DB          DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Host    string `yaml:"host" env:"SERVER_HOST"`
	Port    string `yaml:"port" env:"SERVER_PORT"`
	Timeout int16  `yaml:"timeout"`
}

type AppConfig struct {
	Environment string `yaml:"environment" env:"APP_ENVIRONMENT"`
}

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	Name     string `yaml:"name" env:"DB_NAME"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	SSLMode  string `yaml:"sslMode`
}

func Load() *Config {
	data, fileError := os.ReadFile("./config/config.yml")

	if fileError != nil {
		log.Fatalf("Failed to read config file: %v", fileError)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	if err := applyEnvParameters(&config); err != nil {
		log.Fatalf("Failed to apply environment variables: %v", err)
	}

	return &config
}

func applyEnvParameters(c interface{}) error {
	if c == nil {
		return fmt.Errorf("config cannot be nil")
	}

	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("config must be a pointer")
	}

	v = v.Elem()

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("config must point to a struct")
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			if err := applyEnvParameters(field.Addr().Interface()); err != nil {
				return err
			}

			continue
		}

		envName := fieldType.Tag.Get("env")

		if envName == "" {
			continue
		}

		if envVal := os.Getenv(envName); envVal != "" {
			if field.CanSet() && field.Kind() == reflect.String {
				field.SetString(envVal)
			} else {
				return fmt.Errorf("field %s has unsupported type %s, only string supported", fieldType.Name, field.Kind())
			}
		}
	}

	return nil
}
