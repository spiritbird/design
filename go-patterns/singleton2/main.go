package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var once sync.Once
var singleton *Singleton

func getInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{"singleton"}
	})
	return singleton
}

func main() {
	e1 := getInstance()
	fmt.Println(e1.name)

	e2 := getInstance()
	fmt.Println(e2.name)
}
