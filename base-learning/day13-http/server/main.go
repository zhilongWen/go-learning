package main

import (
	"fmt"
	"net"
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", ":8088")
	tcp, _ := net.ListenTCP("tcp", addr)

	for {

		conn, err := tcp.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	fmt.Println("New connection from", conn.RemoteAddr().String())

	//buf := make([]byte, 1024)
	//str, _ := conn.Read(buf)
	//fmt.Println("Received data:", string(buf[:str]))

	for {
		buf := make([]byte, 1024)
		str, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Received data:", string(buf[:str]))

		conn.Write([]byte("admin-be received your message:" + string(buf[:str])))
	}
}
