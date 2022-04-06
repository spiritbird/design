package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		A int
		B string
	}
	t := T{
		10,
		"testing",
	}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	s.Field(0).SetInt(8)
	s.Field(1).SetString("asc")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
