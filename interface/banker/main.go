package main

import "fmt"

type AbstractBanker interface {
	DoAction()
}

type SaveMoneyBanker struct {
}

func (s SaveMoneyBanker) DoAction() {
	fmt.Println("进行了存款")
}

type TransferBanker struct {
}

func (t TransferBanker) DoAction() {
	fmt.Println("进行了转账")
}

type PayBanker struct {
}

func (p PayBanker) DoAction() {
	fmt.Println("进行了支付")
}

func BankerBusiness(banker AbstractBanker) {
}

func main() {
	sb := &SaveMoneyBanker{}
	sb.DoAction()

	pb := &PayBanker{}
	pb.DoAction()

	tb := &TransferBanker{}
	tb.DoAction()

	BankerBusiness(&SaveMoneyBanker{})
}
