package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := "/Users/topstore/go/src/class1/test.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	//写入数据
	bs := []byte{97, 98, 99, 100} //abcd

	n, err := file.Write(bs[:4])
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	//直接写入字符串

	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)
}
func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
