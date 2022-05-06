package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/topstore/go/src/class1/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bs := make([]byte, 2, 20)
	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
