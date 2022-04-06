package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}
func main() {
	funcValue := reflect.ValueOf(add)

	paramList := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(5)}

	retList := funcValue.Call(paramList)

	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
}
