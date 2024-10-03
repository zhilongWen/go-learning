package main

import (
	"context"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"github.com/go-chassis/openlog"
	"io"
	"os"
)

func main() {

	err := os.Setenv("CHASSIS_CONF_DIR", "/Users/wenzhilong/warehouse/space/go-learning/chassis-demo/_1_quick_start/client/conf")
	if err != nil {
		return
	}

	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	req, err := rest.NewRequest("GET", "http://RESTServer/sayhello/world", nil)
	if err != nil {
		openlog.Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		openlog.Error("do request failed.")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	openlog.Info("REST Server sayhello[GET]: " + string(httputil.ReadBody(resp)))
}
