package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		fmt.Println("error opening file:")
		return
	}

	////z := 0
	//for {
	//	b := make([]byte, 8)
	//	n, err := file.Read(b)
	//
	//	//z++
	//	//if z < 2 {
	//	//	// Seek 每次读完光标往前移动一个字节，所以需要回退一个字节
	//	//	//file.Seek(-1, 1)
	//	//
	//	//
	//	//}
	//
	//
	//
	//	if err != nil {
	//		fmt.Println("error reading file:", err) // error reading file: EOF
	//		return
	//	}
	//
	//	fmt.Println(string(b), n)
	//}

	//  写到文件最前面
	//file.Write([]byte("hello world"))

	////  写到文件最后面
	//file.Seek(0, io.SeekEnd)
	//file.Write([]byte("hello world"))

	//read := bufio.NewReader(file)
	//for {
	//
	//	r, err := read.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("error reading file:", err)
	//		return
	//	}
	//
	//	fmt.Println(r)
	//}

	//all, err := io.ReadAll(file)
	//fmt.Println(string(all))

	//r, err := os.ReadFile("./test.txt")
	//fmt.Println(string(r))

	//dir, err := os.ReadDir("./")
	//fmt.Println(dir)
	//
	//for _, f := range dir {
	//	fmt.Println(f.Type())
	//	fmt.Println(f.Info())
	//	fmt.Println(f.Name())
	//	fmt.Println(f.IsDir())
	//}

	write := bufio.NewWriter(file)
	read := bufio.NewReader(file)

	n := 0
	for {

		n++
		str, err := read.ReadString('\n')

		write.WriteString(strconv.Itoa(n) + " " + str)

		if err != nil {
			fmt.Println("error reading file:", err)
			break
		}
	}

	file.Seek(0, 0)
	write.Flush()

}
