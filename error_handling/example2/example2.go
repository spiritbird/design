package main

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest = errors.New("Bad Request")

	ErrPageMoved = errors.New("Page Moved")
)

func main() {

	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return
		case ErrPageMoved:
			fmt.Println("The Page moved")
			return
		default:
			fmt.Println(err)
			return
		}
	}
}

//webCall performs a web operation.
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrPageMoved
}
