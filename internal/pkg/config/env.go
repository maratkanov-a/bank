package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "BANK"

// Env contains configuration
type Env struct {
	PortHTTP int `envconfig:"PORT_HTTP" default:"8000"`
	PortGRPC int `envconfig:"PORT_GRPC" default:"8443"`

	DatabaseUser            string        `envconfig:"DATABASE_USER" default:"test"`
	DatabasePassword        string        `envconfig:"DATABASE_PASSWORD" default:"test"`
	DatabaseDBName          string        `envconfig:"DATABASE_NAME" default:"bank_test"`
	DatabaseHost            string        `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabasePort            int           `envconfig:"DATABASE_PORT" default:"6432"`
	DatabaseMaxIdleConns    int           `envconfig:"DATABASE_MAX_IDLE_CONNS" default:"10"`
	DatabaseMaxOpenConns    int           `envconfig:"DATABASE_MAX_OPEN_CONNS" default:"10"`
	DatabaseConnMaxLifetime time.Duration `envconfig:"DATABASE_CONN_MAX_LIFETIME" default:"10s"`
}

// GetEnv gets config from env vars
func GetEnv() (*Env, error) {
	var config Env
	err := envconfig.Process(envPrefix, &config)

	return &config, err
}

// GetEnvironmentValues simply returns receiver
func (c *Env) GetEnvironmentValues() *Env {
	return c
}
