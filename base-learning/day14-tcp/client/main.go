package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := new(http.Client)
	//req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req, _ := http.NewRequest(
		"POST",
		"http://localhost:8080/test",
		bytes.NewBuffer([]byte("{\"abc\":111}")),
	)
	req.Header["test"] = []string{"123", "456"}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := res.Body
	data, _ := io.ReadAll(body)
	fmt.Println(string(data))

}
