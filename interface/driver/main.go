package main

import "fmt"

//======>抽象层<======
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层
type BenZ struct {
	//...
}

func (benz *BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
	//...
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhang_3 struct {
	//...
}

func (zhang3 *Zhang_3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

type Li_4 struct {
	//...
}

func (li4 *Li_4) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

func main() {
	//张三开宝马
	//	var bmw Car
	bmw := &Bmw{}

	zhang3 := &Zhang_3{}

	zhang3.Drive(bmw)

	benz := &BenZ{}

	li4 := &Li_4{}

	li4.Drive(benz)
}
