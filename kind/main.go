package main

import (
	"fmt"
	"reflect"
)

type MyInt int64

func ReflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 //指针
	var b MyInt    //自定义类型
	var c rune     //类型别名
	ReflectType(a) //type: kind:ptr  Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
	ReflectType(b) //type:MyInt kind:int64
	ReflectType(c) //type:int31 kind:int32

	type Person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}

	var d = Person{
		name: "wakaka",
		age:  18,
	}

	var e = book{title: "golang"}

	ReflectType(d) //type:Person kind:struct
	ReflectType(e) //type:book kind:struct

}
