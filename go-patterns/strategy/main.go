package main

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

type Addition struct{}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

type Multiplication struct{}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}
func CreateOperation(operator Operator) Operation {
	return Operation{operator}
}

func main() {
	var operationes []Operation
	operatorAdd := CreateOperation(new(Addition))
	operatorMul := CreateOperation(new(Multiplication))

	operationes = append(operationes, operatorAdd)
	operationes = append(operationes, operatorMul)

	for _, operator := range operationes {
		value := operator.Operate(2, 5)
		fmt.Printf("%d\n", value)
	}
}
