package main

import "fmt"

func main() {
	s := student{person{"lisi", 20}, "000"}
	fmt.Println(s.name) //访问结构体的【匿名】字段，student对象没有name字段，访问的是从person继承过来的name字段
	s.run()             //访问结构体的【匿名】字段实现的方法，虽然student对象没有实现run方法，但student继承了person，person类型实现的方法能被直接访问

	t := teacher{person{"wangwu", 25}, "111"}
	fmt.Println(t.p.name) //访问【有名】结构体的字段。不是继承，不能被直接访问，需要指定结构体
	t.p.run()
}

type person struct {
	name string
	age  int
}

// 如果一个struct嵌入另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type student struct {
	person //匿名字段
	number string
}

// 如果一个struct嵌套了另一个【有名】的结构体，叫做组合
type teacher struct {
	p      person //有名字段
	mobile string
}

func (p *person) run() {
	fmt.Println(p.name, " person run")
}

func (s *student) reading() {
	fmt.Println(s.name, " student reading")
}
