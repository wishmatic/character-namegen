package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Host                string `env:"HOST" envDefault:"0.0.0.0"`
	Port                int    `env:"PORT" envDefault:"8080"`
	APIKey              string `env:"API_KEY"`
	LogLevel            string `env:"LOG_LEVEL" envDefault:"info"`
	WriteTimeoutSeconds int    `env:"WRITE_TIMEOUT_SECONDS" envDefault:"600"`
}

func Load() (Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse environment: %w", err)
	}

	return cfg, nil
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
