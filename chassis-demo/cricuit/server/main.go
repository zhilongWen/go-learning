package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/go-chassis/v2/examples/circuit/server/resource"
	"github.com/go-chassis/openlog"
)

func main() {

	// CHASSIS_CONF_DIR=/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/cricuit/server/conf

	println("hello world...")

	chassis.RegisterSchema("rest", &resource.RestFulMessage{}, server.WithSchemaID("RestHelloService"))
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
