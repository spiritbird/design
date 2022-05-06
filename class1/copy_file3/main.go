package main

import (
	"fmt"
	"io/ioutil"
)

func Copy(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("操作失败：", destFile)
		fmt.Println(err)
		return 0, err
	}
	return len(input), nil
}
func main() {
	Copy("/Users/topstore/go/src/class1/test.txt", "/Users/topstore/go/src/class1/test_bak.txt")
}
