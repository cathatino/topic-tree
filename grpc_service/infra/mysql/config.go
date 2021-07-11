package mysql

import (
	"fmt"
	"time"
)

type Config struct {
	Host            string   `toml:"host"`
	Port            string   `toml:"port"`
	User            string   `toml:"user"`
	Password        string   `toml:"password"`
	DbName          string   `toml:"db_name"`
	MaxOpenConns    int      `toml:"max_open_conns"`
	MaxIdleConns    int      `toml:"max_idle_conns"`
	ConnMaxLifetime duration `toml:"conn_max_lifetime"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (c *Config) GetDsn() string {
	return fmt.Sprintf(
		dsnFormat,
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
	)
}

func (c *Config) GetMaxOpenConns() int {
	return c.MaxOpenConns
}

func (c *Config) GetMaxIdleConns() int {
	return c.MaxIdleConns
}

func (c *Config) GetConnMaxLifetime() time.Duration {
	return c.ConnMaxLifetime.Duration
}
