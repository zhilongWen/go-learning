package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("fff")

	var age int
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("age 大于 18 ，age = ", age)
	} else {
		fmt.Println("age 小于 18")
	}

	if age1 := 20; age1 > 18 {
		fmt.Println("age1 大于 18 ，age1 = ", age1)
	} else {
		// else 不能换行
		//else {
		fmt.Println("age1 小于 18")
	}

	if age1 := 20; age1+age > 50 {
		fmt.Println("age1 + age 大于 50 ，age1 + age = ", age1+age)
	} else {
		fmt.Println("age1 + age 小于 50")
	}

	// ===== switch
	var a1 int64 = 10
	//var a2 int32 = 10
	var a2 int64 = 10

	switch a1 {
	case a2, 5, 6, 7: // a2 要与 a1 类型一致
		fmt.Println("ok")
		// 不需要带 break
	default:
		fmt.Println("no ok")
	}

	// for 循环
	for i := 0; i < 10; i++ {
		if i > 5 {
			fmt.Println("interrupt ~~")
			break
		} else {
			fmt.Println(i)
		}
	}

	var str = "wdidnoi是对方空间"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for idx, value := range str {
		fmt.Printf("idx = %d , value = %c \n", idx, value)
	}

	runes := []rune(str)
	fmt.Println(runes)

	// while
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n = ", n)

		if n == 10 {
			continue
		}

		count++

		if n == 99 {
			break
		}
	}
	println(count)

}
