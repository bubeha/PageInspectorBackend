package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server      ServerConfig `json:"server"`
	Application AppConfig    `yaml:"application"`
	Database    DBConfig     `yaml:"database"`
	JWTConfig   `yaml:jwt`
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
	SSLMode  string `yaml:"sslMode"`
}

type JWTConfig struct {
	PublicKeyPath  string        `yaml:"public_path" env:"JWT_PUBLIC_KEY_PATH"`
	PrivateKeyPath string        `yaml:"private_path" env:"JWT_PRIVATE_KEY_PATH"`
	PublicTTL      time.Duration `yaml:"public_ttl" env:"JWT_PUBLIC_TTL"`
	PrivateTTL     time.Duration `yaml:"private_ttl" env:"JWT_PRIVATE_TTL"`
}

func Load() (*Config, error) {
	data, fileError := os.ReadFile("./config/config.yml")

	if fileError != nil {
		return nil, fmt.Errorf("failed to read config file: %v", fileError)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed parsing config file: %v", err)
	}

	if err := applyEnvParameters(&config); err != nil {
		return nil, fmt.Errorf("failed to apply environment variables: %v", err)
	}

	return &config, nil
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

		envVal := os.Getenv(envName)
		if envVal == "" {
			continue
		}

		if !field.CanSet() {
			return fmt.Errorf("field %s cannot be set", fieldType.Name)
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(envVal)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

			if fieldType.Type == reflect.TypeOf(time.Duration(0)) {
				duration, err := time.ParseDuration(envVal)
				if err != nil {
					return fmt.Errorf("invalid duration for field %s: %v", fieldType.Name, err)
				}
				field.SetInt(int64(duration))
			} else {
				intVal, err := strconv.ParseInt(envVal, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid integer for field %s: %v", fieldType.Name, err)
				}
				field.SetInt(intVal)
			}

		default:
			return fmt.Errorf("field %s has unsupported type %s", fieldType.Name, field.Kind())
		}
	}

	return nil
}
