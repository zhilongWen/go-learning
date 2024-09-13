package main

import (
	"fmt"
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
	"os"
)

type HelloResource struct {
}

func (r *HelloResource) SayHi(b *rf.Context) {
	err := b.Write([]byte("hello. go chassis"))
	if err != nil {
		fmt.Println("err = ", err)
	}
	return
}

func (r *HelloResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hello", ResourceFunc: r.SayHi},
	}
}

func main() {

	// CHASSIS_CONF_DIR=/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/hello/server/conf
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("exec getwd failed err = ", err)
		dir = "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo"
	}

	ChassisConfDir := dir + "/hello/server/conf"

	fmt.Println("CHASSIS_CONF_DIR = " + ChassisConfDir)
	err = os.Setenv("CHASSIS_CONF_DIR", ChassisConfDir)
	if err != nil {
		fmt.Println("config CHASSIS_CONF_DIR env path failed. err: ", err)
		return
	}

	println("hello world server...")

	chassis.RegisterSchema("rest", &HelloResource{})

	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	err = chassis.Run()
	fmt.Println(err)
}
