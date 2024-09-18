package main

import "fmt"

func main() {

	fmt.Println("hello world...")

	cat := &Cat{
		Animal: &Animal{
			Name: "Cat",
		},
		sleep: "sleeping",
	}

	cat.Eat()

}

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating \n", a.Name)
}

// Cat 伪继承 animal
type Cat struct {
	*Animal
	sleep string
}
