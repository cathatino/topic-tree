package main

import (
	"fmt"

	"context"
	"log"
	"net"

	"github.com/cathatino/topic-tree/grpc_service/cnf"
	"github.com/cathatino/topic-tree/grpc_service/infra/mysql"
	"github.com/cathatino/topic-tree/grpc_service/internal/modules/user_profile/manager"

	serverPb "github.com/cathatino/topic-tree/grpc_service/pb/server"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	serverPb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *serverPb.HelloRequest) (*serverPb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &serverPb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

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
	//
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serverPb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
