package database

import (
	"context"
	"database/sql/driver"
	"fmt"
	sqlx "github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
)

var (
	defaultDriver = &pq.Driver{}
)

// GetDefaultDriver returns the driver initialized in the init function.
func GetDefaultDriver() driver.Driver {
	return defaultDriver
}

type driverContext struct {
	driver driver.Driver
}

func (d driverContext) OpenConnector(name string) (driver.Connector, error) {
	return &dsnConnector{dsn: name, driver: d.driver}, nil
}

type dsnConnector struct {
	dsn    string
	driver driver.Driver
}

func (d dsnConnector) Connect(_ context.Context) (driver.Conn, error) {
	return d.driver.Open(d.dsn)
}

func (d dsnConnector) Driver() driver.Driver {
	return d.driver
}

// Database contains information about underlying connections
// adds some helpful methods
// and helps to close database connections
type Database struct {
	*sqlx.DB
}

type Options struct {
	User            string
	DBName          string
	Password        string
	Host            string
	Port            int
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// ConnectionString is a connection arguments string representation.
func (c Options) ConnectionString() string {
	return fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		c.User,
		c.DBName,
		c.Password,
		c.Host,
		c.Port,
	)
}

// String is a connection arguments string representation.
func (c Options) String() string {
	return fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		c.User,
		c.DBName,
		"<private ****>",
		c.Host,
		c.Port,
	)
}
