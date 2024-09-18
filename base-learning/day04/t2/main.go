package main

import "fmt"

func main() {

	println("hello world...")

	cat := NewCat("Tom")
	cat.Eat()

	Check(cat)

	cat = NewCat("Tom1")
	cat.Eat()

	Check(cat)
}

// IAnimal =========== start
// IAnimal 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

// IAnimal =========== end

// Animal  =========== start
type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.Name)
}

func NewAnimal(name string) *Animal {
	return &Animal{Name: name}
}

// Animal  =========== end

// Cat   =========== start
type Cat struct {
	*Animal
}

func NewCat(name string) *Cat {
	return &Cat{Animal: NewAnimal(name)}
}

func Check(animal IAnimal) {
	animal.Eat()
}

// Cat  =========== end
