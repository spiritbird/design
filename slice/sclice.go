package main

import (
	"fmt"
	"sync"
)

func main() {
	concurrentAppendSliceNotForceIndex()
	//concurrentAppendSliceForceIndex()
}

//不带索引
func concurrentAppendSliceNotForceIndex() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl = append(sl, num)
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}

//带索引
func concurrentAppendSliceForceIndex() {
	sl := make([]int, 100)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {}(k)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}
