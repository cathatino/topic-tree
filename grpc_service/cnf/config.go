package cnf

import (
	"github.com/BurntSushi/toml"
	"github.com/cathatino/topic-tree/grpc_service/infra/mysql"
)

type Config struct {
	Title    string
	DbConfig mysql.Config `toml:"database"`
}

func LoadConfig(path string) (conf *Config) {
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		panic(err)
	}
	return
}
