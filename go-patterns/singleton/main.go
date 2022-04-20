package main

import (
	"fmt"
	"sync"
)

//Singleton是单例模式对象
type Singleton struct {
	name string
}

var singleton *Singleton

var once sync.Once

//GetInstance 用于获取单例模式对象
func getInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			name: "singleton",
		}
	})
	return singleton
}

func main() {
	e1 := getInstance()
	fmt.Println(e1.name)

	e2 := getInstance()
	fmt.Println(e2.name)
}
