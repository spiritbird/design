package main

import "log"

type MyFunc func(int) int

func LogDecorate(fn MyFunc) (resultFunc MyFunc) {

	resultFunc = func(n int) (finalResult int) {
		log.Println("Starting the execution with integer", n)

		finalResult = fn(n)

		log.Println("Execution is completed with the result", finalResult)
		return
	}
	return
}

type Component interface {
	Double() int
}

type ConcreteComponent struct {
	Num1 int
}

func (c *ConcreteComponent) Double() int {
	return c.Num1 * 2
}

//匿名组合组合方式
type MulDecorator struct {
	Component
	num int
}

//WarpMulDecorator 返回的是接口类型Component
func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

//Double MulDecorator
func (d *MulDecorator) Double() int {
	return d.Component.Double() * d.num
}

// AddDecorator 匿名组合组合方式
type AddDecorator struct {
	Component
	num int
}

// WarpAddDecorator 返回的是接口类型Component
func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c,
		num,
	}
}

// Double AddDecorator实现了Component接口
func (d *AddDecorator) Double() int {
	return d.Component.Double() + d.num
}

func main() {
	var fn MyFunc
	fn = func(n int) int {
		return n * 2
	}
	resultFunc := LogDecorate(fn)
	log.Printf("result:%d\n", resultFunc(5))
	log.Println("匿名组合实现装饰模式")
	c := &ConcreteComponent{
		Num1: 1,
	}

	ww := WarpMulDecorator(c, 3)

	log.Printf("匿名组合实现装饰模式,result:%d\n", ww.Double())

	wd := WarpMulDecorator(c, 3)

	log.Printf("匿名组合实现装饰模式,result:%d\n", wd.Double())
}
