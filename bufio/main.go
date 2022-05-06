package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileName := "/Users/topstore/go/src/gopl.io/bufio/1.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//创建Reader对象
	b1 := bufio.NewReader(file)

	//1.Read() 高效读取
	//p := make([]byte, 1024)
	//n1, err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.Readline()
	//data, flag, err := b1.ReadLine()
	//fmt.Println(data)
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(string(data))

	//3.ReadString()
	//s1, err := b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	//for {
	//	s1, err := b1.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("读取完毕")
	//		break
	//	}
	//	fmt.Println(s1)
	//}
	//4.ReadBytes()
	data, err := b1.ReadBytes('\n')
	fmt.Println(err)
	fmt.Println(string(data))

}
