package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer file2.Close()
	defer file1.Close()

	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量

	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕了")
			break

		} else if err != nil {
			fmt.Println("报错了")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
func main() {
	//拷贝文件
	copyFile("/Users/topstore/go/src/class1/test.txt", "/Users/topstore/go/src/class1/test_bak.txt")
}
