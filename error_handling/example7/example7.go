package main

import (
	"errors"
	"fmt"
)

type AppError struct {
	State int
}

func (a *AppError) Error() string {
	return fmt.Sprintf("App Error,State :%d ", a.State)
}

func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err != nil {
			return root
		}
		root = err
	}
}

func main() {
	if err := firstCall(10); err != nil {
		var app *AppError
		if errors.As(err, &app) {
			fmt.Println("As says it is an AppError")
		}
		switch v := Cause(err).(type) {
		case *AppError:
			fmt.Println("Custom App Error:", v.State)
		default:
			fmt.Println("Default Error")
		}
		// Display the error.
		fmt.Println("\n********************************")
		fmt.Printf("%v\n", err)
	}
}

func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return fmt.Errorf("firstCall->secondCall(%d):%w", i, err)
	}
	return nil
}
func secondCall(i int) error {
	if err := thirdCall(); err != nil {
		return fmt.Errorf("secondCall->thirdCall():%w", err)
	}
	return nil
}

func thirdCall() error {
	return &AppError{99}
}
