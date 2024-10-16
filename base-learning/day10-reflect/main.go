package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

func (u User) SayHello(name string) {
	fmt.Println("我的名字 ", name)
}

type Student struct {
	Course string
	User
}

func (s *Student) GetCourse(course string) {
	fmt.Println("s的课程 ", course)
}

// 方法名第一个字母小写表示私有的，其他文件种不能调用，第一个字母大写表示公有的，其他文件可以调用
func main() {

	u := User{Name: "小明", Age: 20, Sex: true}
	//Check(u)
	//
	//println("============")
	//
	//s := Student{"11", User{}}
	//Check(s)

	//Check(u)

	s := Student{"11", u}
	Check(&s)

}

func Check(inter interface{}) {
	//inter.(User).SayHello("小明") // inter.(User) 类型断言，将v转换为User类型，类似强转

	//switch inter.(type) {
	//case User:
	//	fmt.Println("v是User类型")
	//	fmt.Println(inter.(User).Name)
	//case Student:
	//	fmt.Println("v是Student类型")
	//	fmt.Println(inter.(Student).Course)
	//}

	// TypeOf() 获取类型信息，ValueOf() 获取值信息
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	//
	fmt.Println(t)
	fmt.Println(v)
	//
	//// NumField() 获取结构体字段数量 Field() 获取结构体字段值
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println(v.Field(i))
	//}

	//// FieldByIndex() 获取结构体字段值，通过索引获取,FieldByName() 获取结构体字段值，通过字段名获取
	//fmt.Println(v.FieldByIndex([]int{0}))
	//fmt.Println(v.FieldByIndex([]int{1}))
	////fmt.Println(v.FieldByIndex([]int{2})) // panic: reflect: Field index out of range
	//
	//fmt.Println(v.FieldByName("Course"))
	//fmt.Println(v.FieldByIndex([]int{1, 0}))
	//fmt.Println(v.FieldByIndex([]int{1, 1}))
	//fmt.Println(v.FieldByIndex([]int{1, 2}))

	//ty := t.Kind()
	//
	//if ty == reflect.Struct {
	//	fmt.Println("v是结构体类型")
	//}
	//
	//if ty == reflect.String {
	//	fmt.Println("v是字符串类型")
	//}

	//e := v.Elem()
	//e.FieldByName("Name").SetString("小红")
	//fmt.Println(inter)

	fmt.Println(v.Method(0))
	m := v.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("120")})
}
