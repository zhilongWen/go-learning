package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Request struct {
	Num0 int
	Num1 int
}

type Response struct {
	Result int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8088")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := Request{10, 20}
	var res Response

	// Call 同步调用
	//err = client.Call("Server.Add", req, &res)

	// Go 异步调用
	call := client.Go("Server.Add", req, &res, nil)
	fmt.Println("我还可以干其他的...")

	<-call.Done // 等待调用结束

	fmt.Println(res.Result)
}
