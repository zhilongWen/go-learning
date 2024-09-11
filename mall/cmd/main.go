package main

import (
	"mall/conf"
	"mall/routes"
)

func main() {

	conf.Init()

	r := routes.Router()
	_ = r.Run(conf.HttpPort)

}
