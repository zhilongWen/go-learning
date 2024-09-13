package main

import (
	"context"
	"fmt"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"log"
	"net/http"
	"os"

	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type SimpleResource struct {
}

func (r *SimpleResource) SayHi(b *rf.Context) {
	req, _ := rest.NewRequest(http.MethodGet, "http://HelloServer/hello", nil)
	restInvoker := core.NewRestInvoker()
	resp, err := restInvoker.ContextDo(context.TODO(), req)
	if err != nil {
		log.Println(err)
		return
	}

	b.Write(httputil.ReadBody(resp))

	fmt.Println("ddddd")

	return
}

func (r *SimpleResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hi", ResourceFunc: r.SayHi},
	}
}

func main() {

	//  CHASSIS_CONF_DIR=/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/hello/client/conf
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("exec getwd failed err = ", err)
		dir = "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo"
	}

	ChassisConfDir := dir + "/hello/client/conf"

	fmt.Println("CHASSIS_CONF_DIR = " + ChassisConfDir)
	err = os.Setenv("CHASSIS_CONF_DIR", ChassisConfDir)
	if err != nil {
		fmt.Println("config CHASSIS_CONF_DIR env path failed. err: ", err)
		return
	}

	println("hello world client...")

	chassis.RegisterSchema("rest", &SimpleResource{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}

	err = chassis.Run()
	fmt.Println(err)
}
