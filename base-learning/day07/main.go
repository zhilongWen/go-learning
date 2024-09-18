package main

import (
	"fmt"
	"reflect"
)

func main() {

	println("hello world...")

	var x float64 = 3.14159
	fmt.Println("type : ", reflect.TypeOf(x)) // type :  float64

	p := &Person{Name: "tom"}
	value := reflect.ValueOf(p)
	method := value.MethodByName("Talk")

	method.Call(nil)
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi , my name is ", p.Name)
}
