package main

import (
	"chassis-demo/v1_rest/server/schema"
	io_grpc_examples_helloworld "chassis-demo/v2_grpc/helloworld/io.grpc.examples.helloworld"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"log"
	"os"
)

func main() {

	// CHASSIS_CONF_DIR=/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/v2_grpc/server/conf
	_ = os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/v2_grpc/server/conf")

	println("grpc ...")

	// [register]
	//1. Registration Service
	chassis.RegisterSchema("grpc", &schema.Server{}, server.WithRPCServiceDesc(&io_grpc_examples_helloworld.Greeter_ServiceDesc))
	// [register]

	// [init]
	//2. Initialize Service
	if err := chassis.Init(); err != nil {
		log.Println("chassis init failed:", err)
		return
	}

	//3. Start service, blocking waiting for service to roll out
	if err := chassis.Run(); err != nil {
		log.Println("chassis run failed:", err)
		return
	}
	// [init]
}
