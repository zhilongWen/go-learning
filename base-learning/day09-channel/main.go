package main

import (
	"fmt"
	"sync"
)

func main() {

	//Run()

	//go Run() // 没有打印，新开了一个协程，与 main 所在的协程无关
	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//}

	var wg sync.WaitGroup
	wg.Add(1)
	go Run(&wg)
	wg.Wait()

	println("main end.")

}

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了...")
	wg.Done()
}
