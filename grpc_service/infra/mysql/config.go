package mysql

import (
	"fmt"
	"time"
)

type Config struct {
	host            string
	port            string
	user            string
	password        string
	dbName          string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}

func (c *Config) GetDsn() string {
	return fmt.Sprintf(
		dsnFormat,
		c.user,
		c.password,
		c.host,
		c.port,
		c.dbName,
	)
}

func (c *Config) GetMaxOpenConns() int {
	return c.maxOpenConns
}

func (c *Config) GetMaxIdleConns() int {
	return c.maxIdleConns
}

func (c *Config) GetConnMaxLifetime() time.Duration {
	return c.connMaxLifetime
}
