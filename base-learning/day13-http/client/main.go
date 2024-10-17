package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", ":8088")
	conn, _ := net.DialTCP("tcp", nil, addr)

	//conn.Write([]byte("Hello, world!"))

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		_, _ = conn.Write(data)

		rb := make([]byte, 1024)
		_, _ = conn.Read(rb)
		fmt.Println(string(rb))
	}
}
