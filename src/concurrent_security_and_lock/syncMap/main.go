package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	m  sync.Map
)

// go 原有的map不支持并发安全，如果同时多个并发对同一个map操作，就会导致 concurrent map writes
//var m = make(map[int]int)

func get(key int) (any, bool) {
	return m.Load(key)
}

func set(key, value int) {
	m.Store(key, value)
}

func main() {
	// sync.Map 并发安全的map
	for i := 0; i < 20; i++ {
		wg.Add(1)
		i := i
		go func() {
			set(i, i+100)
			value, _ := get(i)
			fmt.Printf("key: %v, value: %v \n", i, value)
			wg.Done()
		}()
	}
	wg.Wait()

}
