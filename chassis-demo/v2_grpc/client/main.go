package main

import (
	"chassis-demo/v2_grpc/helloworld/helloworld"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/client"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/core/common"
	"github.com/go-chassis/openlog"
	"os"
)

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/grpc/client/
func main() {

	_ = os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/v2_grpc/client/conf")

	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}

	//declare reply struct
	reply := &helloworld.HelloReply{}
	//header will transport to target service
	ctx := common.NewContext(map[string]string{
		"X-User": "pete",
	})
	invoker := core.NewRPCInvoker()

	//Invoke with micro service name, schemas ID and operation ID
	err := invoker.Invoke(
		ctx,
		"v2_grp_server",
		"helloworld.Greeter",
		"SayHello",
		&helloworld.HelloRequest{Name: "Peter"},
		reply,
		core.WithProtocol("grpc"),
	)
	if err != nil {
		openlog.Error("error" + err.Error())
	}
	openlog.Info(reply.Message)
}
