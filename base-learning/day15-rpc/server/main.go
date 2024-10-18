package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Request struct {
	Num0 int
	Num1 int
}

type Response struct {
	Result int
}

func (s *Server) Add(req *Request, res *Response) error {
	fmt.Println("我要睡 5 秒")
	time.Sleep(5 * time.Second)
	res.Result = req.Num0 + req.Num1
	return nil
}

func main() {

	_ = rpc.Register(&Server{})
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("ListenTCP error:", err)
		return
	}

	err = http.Serve(listen, nil)
	if err != nil {
		fmt.Println("Serve error:", err)
		return
	}
}
