package main

import (
	"fmt"

	"github.com/cathatino/topic-tree/grpc_service/cnf"
	"github.com/cathatino/topic-tree/grpc_service/infra/mysql"
)

func main() {
	config := cnf.LoadConfig("../cnf/test.toml")
	engine, err := mysql.NewEngine(config.DbConfig)
	fmt.Println(engine, err)
}
