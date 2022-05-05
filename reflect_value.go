package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32 ,value is %f\n", v.Float())
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", v.Float())

	}
}
func main() {
	var s float32 = 3.14
	var b int64 = 100
	reflectValue(s)
	reflectValue(b)
}
