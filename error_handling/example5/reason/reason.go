package main

import (
	"fmt"
	"log"
)

type customError struct {
}

func (c *customError) Error() string {
	return "Find the bug."
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	fmt.Printf("Type of value strored inside the interface:%T\n", err)

	if _, err = fail(); err != nil {
		fmt.Printf("Type of value stored inside the interface:%T\n", err)
	}
	log.Println("No Error")
}
