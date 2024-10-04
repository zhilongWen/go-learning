package main

import (
	"chassis-demo/_2_quick_start/schemas/helloworld"
	"context"
	_ "errors"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/client"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core"
	_ "github.com/go-chassis/go-chassis/v2/core/lager"
	"os"
)

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {

	err := os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/_2_quick_start/client/conf")
	if err != nil {
		return
	}

	//Init framework
	if err := chassis.Init(); err != nil {
		//lager.Logger.Error("Init failed." + err.Error())
		return
	}
	//declare reply struct
	reply := &helloworld.HelloReply{}
	//Invoke with microservice name, schema ID and operation ID
	if err := core.NewRPCInvoker().Invoke(context.Background(), "RPCServer", "helloworld.Greeter", "SayHello",
		&helloworld.HelloRequest{Name: "Peter"}, reply, core.WithProtocol("grpc")); err != nil {
		//lager.Logger.Error("error" + err.Error())
	}
	//lager.Logger.Info(reply.Message)
}
