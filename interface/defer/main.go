package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}

func DeferFunc11(i int) (t int) {

	fmt.Println("t = ", t)

	return 2
}
func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

	defer func() {
		t = t * 10
	}()

	return 1
}
func main() {
	returnAndDefer()
	DeferFunc11(10)
	fmt.Println(returnButDefer())
}
