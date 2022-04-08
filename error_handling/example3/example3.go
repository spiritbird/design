package main

import (
	"fmt"
	"reflect"
)

//An UnmarshalTypeError describes a jSON value that was
//not appropriate for a value of a specific Go type

type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

//Error implements the error interface
func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type" + e.Type.String()
}

//An InvalidUnmarshalError describe an invalid argument passed to Unmarshal;
//(The argument to Unmarshal must be a noo-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

//Error implements the error interface
func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json:Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json:Unmarshal(nil " + e.Type.String() + ")"
}

//user is a type for use in the Unmarshall call.
type user struct {
	Name string
}

//Unmarshal simulates an unmarshal call that always fails
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	s := reflect.TypeOf(v)
	fmt.Println(s)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}

func main() {

	var u user
	err := Unmarshal([]byte(`{"name":"bill"}`), &u)
	if err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError:Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError:Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		return
	}
}
