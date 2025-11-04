package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port        string
	Environment string
	Timeout     int16
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:        ":8080",
			Environment: "development",
			Timeout:     60,
		},
	}
}
