package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type AppError struct {
	State int
}

func (c *AppError) Error() string {
	return fmt.Sprintf("App Error,State: %d", c.State)
}

func main() {
	if err := firstCall(10); err != nil {
		switch v := errors.Cause(err).(type) {
		case *AppError:
			fmt.Println("Custom App Error:", v.State)
		default:
			fmt.Println("Default Error")
		}
		// Display the stack trace for the error.
		fmt.Println("\nStack Trace\n********************************")
		fmt.Printf("%+v\n", err)
		fmt.Println("\nNo Trace\n********************************")
		fmt.Printf("%v\n", err)
	}

}

func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return errors.Wrapf(err, "firstCall->secondCall(%d)", i)
	}
	return nil
}
func secondCall(i int) error {
	if err := thirdCall(i); err != nil {
		return errors.Wrapf(err, "secondCall->thirdCall(%d)", i)
	}
	return nil
}
func thirdCall(i int) error {
	return &AppError{i}
}
