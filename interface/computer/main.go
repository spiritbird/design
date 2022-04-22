package main

import "fmt"

type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type CPU interface {
	Calculate()
}

type Computer struct {
	cpu  CPU
	mem  Memory
	card Card
}

func (c *Computer) DoWork() {
	c.mem.Storage()
	c.card.Display()
	c.cpu.Calculate()
}

func NewComputer(cpu CPU, mem Memory, card Card) *Computer {
	return &Computer{
		cpu:  cpu,
		mem:  mem,
		card: card,
	}
}

type IntelCPU struct {
	CPU
}

func (i *IntelCPU) Calculate() {
	fmt.Println("Intel CPU开始计算了...")
}

type IntelMemory struct {
	Memory
}

func (i IntelMemory) Storage() {
	fmt.Println("Intel Memory 开始储存了...")
}

type IntelCard struct {
	Card
}

func (i IntelCard) Display() {
	fmt.Println("Intel Card 开始显示了...")
}

type KingstonMemory struct {
	Memory
}

func (k KingstonMemory) Storage() {
	fmt.Println("Kingston Memory storage...")
}

type NvidiaCard struct {
	Card
}

func (n NvidiaCard) Display() {
	fmt.Println("Nvidia Card Display ...")
}

func main() {
	//intel系列的电脑
	com1 := NewComputer(&IntelCPU{}, &IntelMemory{}, &IntelCard{})
	com1.DoWork()

	//杂牌子
	com2 := NewComputer(&IntelCPU{}, &KingstonMemory{}, &NvidiaCard{})
	com2.DoWork()
}
