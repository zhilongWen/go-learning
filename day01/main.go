package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	n1 := 1000

	fmt.Printf(" n1 = %T \n", n1)
	fmt.Println(unsafe.Sizeof(n1))

	var c1 byte
	c1 = '0'
	fmt.Println(c1)
	fmt.Printf("c1 = %c \n", c1)

	var c2 int = '北'
	fmt.Println(c2)
	fmt.Printf("c2 = %c \n", c2)

	// overflows
	// 如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]直接可以保存到 byte
	// 如果我们保存的字符对应码值大于 255,这时我们可以考虑使用 int 类型保存
	//var c3 byte = '北'
	//fmt.Println(c3)
	//fmt.Printf("c3 = %c \n", c3)

	var c4 int = 22269
	fmt.Printf("c4 = %c \n", c4)

	var c5 = 10 + 'a'
	fmt.Println(c5) // 10 + 97 = 107

	var c6 bool = true
	fmt.Printf("c6 = %T \n", c6)
	fmt.Println(unsafe.Sizeof(c6))

	var c7 string = "asdkl 说的那可能是"
	fmt.Println(c7)
	fmt.Println(c7[0])
	fmt.Printf("idx0 = %c \n", c7[0])
	//c7[0] = 'l' // 在 Go 中字符串是不可变的

	c8 := "abc\nbv"
	fmt.Println(c8)

	// 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	c9 := `abc\nbv`
	fmt.Println(c9)

	var i int32 = 1000
	var i1 float64 = float64(i)
	fmt.Printf("i = %v i1 = %v \n", i, i1)
	fmt.Printf("i = %T i1 = %T \n", i, i1)

	// 大范围转小范围结果溢出
	// i2 = -24 i2 = int8
	var i2 int8 = int8(i)
	fmt.Printf("i2 = %v i2 = %T \n", i2, i2)

	var s1 int = 99
	var s2 float64 = 23.00001
	var s3 bool = true
	var s4 byte = 'h'
	var ss string

	ss = fmt.Sprintf("%d", s1)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = fmt.Sprintf("%f", s2)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = fmt.Sprintf("%t", s3)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = fmt.Sprintf("%c", s4)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = strconv.FormatInt(int64(s1), 10)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = strconv.FormatFloat(s2, 'f', 10, 64)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = strconv.FormatBool(s3)
	fmt.Printf("ss type = %T str = %q \n", ss, ss)

	ss = "true1"
	// 返回两个值
	parseBool, err := strconv.ParseBool(ss)
	fmt.Printf("parseBool type = %T parseBool = %v \n", parseBool, parseBool)
	fmt.Printf("err type = %T err = %v \n", err, err)

	// 不想获取异常值使用 _
	parseBool1, _ := strconv.ParseBool(ss)
	fmt.Printf("parseBool1 type = %T parseBool1 = %v \n", parseBool1, parseBool1)

	// 指针
	fmt.Println("ss 的地址 = ", &ss)

	var d1 int = 10
	fmt.Println("d1 的地址 = ", &d1)

	// d1_ptr 是一个指针变量
	// d1_ptr 的类型是 *int
	// d1_ptr 本身的值是 &d1
	var d1_ptr *int = &d1
	fmt.Printf("d1_ptr 的值 %v \n", d1_ptr)
	fmt.Printf("d1_ptr 的地址 %v \n", &d1_ptr)
	fmt.Printf("d1_ptr 指向的值 %v \n", *d1_ptr)
	// 通过地址改变 d1 的值
	*d1_ptr = 100
	fmt.Printf("d1 = %v \n", *d1_ptr)

	// ================

	// 只保留整数部分
	var a1 float64 = 10 / 4
	fmt.Println(a1)

	fmt.Println(10.0 / 4)

	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("-10%-3=", -10%-3) // -1

}
