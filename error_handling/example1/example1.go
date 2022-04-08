package main

import "fmt"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}
func main() {
	if err := WebCall(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("life is good")
}

func WebCall() error {
	return New("Bad Request")
}
