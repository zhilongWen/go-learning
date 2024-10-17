package main

import (
	"fmt"
	"sync"
	"time"
)

func SYncClass() {

	// 依次调用lockFun()函数，模拟多个协程同时执行
	//lockFun()
	//lockFun()
	//lockFun()
	//lockFun()

	// 并发执行lockFun()函数, time.Sleep() 没有作用，因为协程是并发执行的
	//go lockFun()
	//go lockFun()
	//go lockFun()
	//go lockFun()
	/**
	func lockFun() {
		fmt.Println("疯狂刮痧....")
		time.Sleep(1 * time.Second)
	}
	*/

	// 使用sync.Mutex 锁进行同步
	//lock := &sync.Mutex{}
	//
	//go lockFun(lock)
	//go lockFun(lock)
	//go lockFun(lock)
	//go lockFun(lock)
	/**
	func lockFun(lock *sync.Mutex) {
		lock.Lock()
		fmt.Println("疯狂刮痧....")
		time.Sleep(1 * time.Second)
		lock.Unlock()
	}
	*/

	// RWMutex 读不不互斥，有读时不能写，写时写读不能操作
	//lock := &sync.RWMutex{}
	//
	//go lockFun(lock)
	//go lockFun(lock)
	//
	//go rwRLockFun(lock)
	//go rwRLockFun(lock)
	//go rwRLockFun(lock)
	//go rwRLockFun(lock)

	// sync.Once 只执行一次函数
	//once := &sync.Once{}
	//for i := 0; i < 10; i++ {
	//	once.Do(func() {
	//		fmt.Println(i)
	//	})
	//}
	//for i := 0; i < 10; i++ {
	//	once.Do(func() {
	//		fmt.Println(i)
	//	})
	//}

	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	time.Sleep(8 * time.Second)
	//	fmt.Println(8)
	//	wg.Done()
	//}()
	//
	//go func() {
	//	time.Sleep(6 * time.Second)
	//	fmt.Println(6)
	//	wg.Done()
	//}()
	//
	//// 可一看到 6 和 8 打印
	////time.Sleep(20 * time.Second)
	//// 无法看到打印，在实际情况中等待时间时不确定的
	////time.Sleep(3 * time.Second)
	//// 使用 WaitGroup 等待协程执行完毕
	//wg.Wait()

	// fatal error: concurrent map read and map write
	//m := make(map[int]int)
	//m := &sync.Map{}
	//
	//go func() {
	//	for {
	//		//m[1] = 1
	//		m.Store(1, 1)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		//fmt.Println(m[1])
	//		fmt.Println(m.Load(1))
	//	}
	//}()

	//pool := &sync.Pool{}
	//pool.Put(1)
	//pool.Put(2)
	//pool.Put(3)
	//pool.Put(4)
	//pool.Put(5)
	//pool.Put(6)
	//pool.Put("a")
	//
	//for i := 0; i < 10; i++ {
	//	time.Sleep(1 * time.Second)
	//	// 随机获取一个值
	//	fmt.Println(pool.Get())
	//}

	co := sync.NewCond(&sync.Mutex{})
	go func() {
		co.L.Lock()
		fmt.Println("lock 1")
		co.Wait()
		fmt.Println("unlock 1")
		co.L.Unlock()
	}()
	go func() {
		co.L.Lock()
		fmt.Println("lock 2")
		co.Wait()
		fmt.Println("unlock 2")
		co.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	co.Broadcast()
	time.Sleep(2 * time.Second)

}

func lockFun(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("疯狂刮痧....")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

func rwRLockFun(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("疯狂治疗....")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}

func main() {
	SYncClass()
	//for {
	//}
}
