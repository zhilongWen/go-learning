package main

import (
	pb "chassis-demo/_2_quick_start/schemas/helloworld"
	"context"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/openlog"
	"os"
)

type Server struct {

	// 确保嵌入未实现的方法，必须要写的否则报错
	// FATAL: [core] grpc: Server.RegisterService found the handler of type *main.Server that does not satisfy helloworld.GreeterServer
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// https://go-chassis.readthedocs.io/en/latest/getstarted/writing-rpc.html
func main() {

	err := os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/_2_quick_start/server/conf")
	if err != nil {
		return
	}
	chassis.RegisterSchema("grpc", &Server{}, server.WithRPCServiceDesc(&pb.Greeter_ServiceDesc))
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed.")
		return
	}
	err = chassis.Run()
	if err != nil {
		return
	}
}
