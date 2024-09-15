package main

import (
	"chassis-demo/v1_rest/server/schema"
	"fmt"
	"github.com/go-chassis/go-chassis/v2"
	"os"
)

func main() {

	// http://127.0.0.1:9000/sayhello

	// CHASSIS_CONF_DIR=/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/v1_rest/server/conf

	_ = os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/v1_rest/server/conf")

	println("hello world go rest...")

	confPath, err := os.Getwd()
	if err != nil {
		fmt.Println("error get conf path, ", err)
		return
	}
	fmt.Println("conf path: ", confPath)

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	println("env: ", env)

	//1. Register Service Processing Logic
	chassis.RegisterSchema("rest", &schema.Server{})

	//2. Initialize servier
	err = chassis.Init()
	if err := err; err != nil {
		fmt.Println("chassis init failed:", err)
		return
	}

	//3. start service
	err = chassis.Run()
	if err != nil {
		fmt.Println("chassis run failed:", err)
		return
	}
}
