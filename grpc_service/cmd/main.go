package main

import (
	"fmt"

	"github.com/cathatino/topic-tree/grpc_service/cnf"
	"github.com/cathatino/topic-tree/grpc_service/infra/mysql"
	"github.com/cathatino/topic-tree/grpc_service/internal/modules/user_profile/manager"
)

func main() {
	config := cnf.LoadConfig("../cnf/test.toml")
	engine, _ := mysql.NewEngine(config.DbConfig)
	userManager := manager.NewUserManager(engine, engine)
	userModel, err := userManager.CreateUser(
		"auth_method",
		"profile",
		"meta_data",
	)
	fmt.Println(userModel, err)
}
