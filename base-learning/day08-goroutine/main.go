package main

func main() {

	/*
			channel 分为 5 种
				1.可读可取 c := make(chan int)
				2.可读 var readCh <- chan int = c
				3.可写 var writeCh chan <- int = c
				4.有缓冲 c := make(chan int,10)
		        5.无缓冲 c := make(chan int)
	*/

	//// 有缓冲，必须先存后取否则 fatal error: all goroutines are asleep - deadlock!
	c1 := make(chan int, 5)
	//c1 <- 1                  // 存
	//fmt.Println("c1:", <-c1) // 读
	////fmt.Println("c1:", <-c1) // 读 ，存一个只能取一个 fatal error: all goroutines are asleep - deadlock!

	// 无缓冲
	//c2 := make(chan int)
	////c2 <- 1 // 没有地方存，必须先定义缓冲，否则 fatal error: all goroutines are asleep - deadlock!
	//go func() {
	//	fmt.Println("我先执行了。。。")
	//	c2 <- 1
	//}()
	//fmt.Println("c2:", <-c2)
	//fmt.Println("c2:", <-c2)  // 同样存一个只能取一个

	//// 无缓冲区，一存一取，同步
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("c2 write :", i)
	//		c2 <- i
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println("c2 read :", <-c2) // 一定是写了，才能读到
	//}
	//
	//fmt.Println("===================")
	//
	//// 有缓冲区，先存满，再取，之后只要缓冲没有满就可以存，所以存和取异步，但是必须缓冲有才能取
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		fmt.Println("c1 write :", i)
	//		c1 <- i
	//	}
	//}()
	//
	//for i := 0; i < 100; i++ {
	//	fmt.Println("c1 read :", <-c1) // 一定是写了，才能读到
	//}

	//var readCh <-chan int = c1 // readCh  为只可读的 c1
	////readCh <- 1  // Invalid operation: readCh <- 1 (send to the receive-only type <-chan int)
	//
	//var writeCh chan<- int = c1 // writeCh 为只可写的 c1
	////<- writeCh // Invalid operation: <- writeCh (receive from the send-only type chan<- int)
	//
	//writeCh <- 1
	//writeCh <- 2
	//fmt.Println(<-readCh) // 1
	//fmt.Println(<-c1)     // 2

	c1 <- 1
	c1 <- 2
	c1 <- 3
	//close(c1) // panic: send on closed channel
	c1 <- 4
	//close(c1) // 关闭 channel，使得读者读不到数据，但是写者依然可以往里面写数据
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)

	//close(c1) // 如果要遍历 channel 必须先关闭，否则 fatal error: all goroutines are asleep - deadlock!
	//for v := range c1 {
	//	fmt.Println(v)
	//}

	// 一个 case
	var c = make(chan int)
	var readCh <-chan int = c
	var writeCh chan<- int = c

	go SetChannel(writeCh)

	GetChannel(readCh)

}

func GetChannel(readCh <-chan int) {

	for i := 0; i < 10; i++ {
		println("read : ", <-readCh)
	}

	// 不行？？？
	//Waiting:
	//	for i := 0; i < 10; i++ {
	//		select {
	//		case <-readCh:
	//			goto Read
	//		default:
	//			//fmt.Println("waiting ", i)
	//			time.Sleep(10)
	//		}
	//
	//	}
	//
	//Read:
	//	println("read : ", <-readCh)
	//	goto Waiting

}

func SetChannel(writeCh chan<- int) {
	for i := 0; i < 10; i++ {
		writeCh <- i
	}
}
