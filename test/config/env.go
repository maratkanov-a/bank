package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "TEST_BANK"

// Config contains configuration for clients in integration tests
type Config struct {
	AppAddress              string        `split_words:"true" default:"localhost:8843"`
	DatabaseUser            string        `envconfig:"DATABASE_USER" default:"test"`
	DatabasePassword        string        `envconfig:"DATABASE_PASSWORD" default:"test"`
	DatabaseDBName          string        `envconfig:"DATABASE_NAME" default:"bank_test"`
	DatabaseHost            string        `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabasePort            int           `envconfig:"DATABASE_PORT" default:"6432"`
	DatabaseMaxIdleConns    int           `envconfig:"DATABASE_MAX_IDLE_CONNS" default:"10"`
	DatabaseMaxOpenConns    int           `envconfig:"DATABASE_MAX_OPEN_CONNS" default:"10"`
	DatabaseConnMaxLifetime time.Duration `envconfig:"DATABASE_CONN_MAX_LIFETIME" default:"10s"`
}

// FromEnv gets config from env vars
func FromEnv() (*Config, error) {
	config := &Config{}
	err := envconfig.Process(envPrefix, config)
	return config, err
}
