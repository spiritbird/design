package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
func main() {
	Copy("/Users/topstore/go/src/class1/test.txt", "/Users/topstore/go/src/class1/test_bak.txt")
}
