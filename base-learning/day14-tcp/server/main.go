package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("GET request received"))
	case "POST":

		// header 中 key 的第一个字母会大写
		//fmt.Println("req.Header test: ", req.Header["test"])
		fmt.Println("req.Header test: ", req.Header["Test"])

		b, _ := io.ReadAll(req.Body)
		header := res.Header()
		header.Set("test", "ttt")
		header["abcdd"] = []string{"123", "456"}

		res.Write([]byte("POST request received " + string(b)))
	default:
		res.Write([]byte("Method not allowed"))
	}
}

func main() {

	// http://localhost:8080/test
	//http.HandleFunc("/test", handler)
	//http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}
