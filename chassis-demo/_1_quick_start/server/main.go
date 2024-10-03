package main

import (
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
	"os"
)

type RestFulHello struct{}

func (r *RestFulHello) SayHello(b *rf.Context) {
	err := b.Write([]byte("get user id: " + b.ReadPathParameter("userid")))
	if err != nil {
		return
	}
}

func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{
			Method:       http.MethodGet,
			Path:         "/sayhello/{userid}",
			ResourceFunc: r.SayHello,
			Returns:      []*rf.Returns{{Code: 200}},
		},
	}
}

// https://go-chassis.readthedocs.io/en/latest/getstarted/writing-rest.html
func main() {

	// http://127.0.0.1:5001/sayhello/world

	err := os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/_1_quick_start/server/conf")
	if err != nil {
		return
	}

	// register struct
	chassis.RegisterSchema("rest", &RestFulHello{})
	//start all server you register in server/schemas.
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed. " + err.Error())
		return
	}
	err = chassis.Run()
	if err != nil {
		return
	}
}
