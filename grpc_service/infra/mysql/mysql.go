package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewEngine(config Config) (
	engine *xorm.Engine,
	err error,
) {
	engine, err = xorm.NewEngine(mysqlDriverName, config.GetDsn())
	if err == nil {
		return
	}
	engine.SetConnMaxLifetime(config.GetConnMaxLifetime())
	engine.SetMaxIdleConns(config.GetMaxIdleConns())
	engine.SetMaxOpenConns(config.GetMaxOpenConns())
	return
}
