package main

import "sync"

type Singleton struct {
	name string
}

var once sync.Once

var singleton *Singleton

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
	println(e1.name)

	e2 := getInstance()
	println(e2.name)
}
