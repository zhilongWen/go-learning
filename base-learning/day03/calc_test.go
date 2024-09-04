package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {

	res := addUpper(10)

	if res > 10 {
		fmt.Println(res)
	} else {
		fmt.Println("err")
	}

}
