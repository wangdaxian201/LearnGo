package main

import (
	"fmt"
	"sync"
)

// 并发安全和锁

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex
	relock sync.RWMutex
)

func add() {
	lock.Lock()
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	lock.Unlock()
	wg.Done()
}

func add1() {
	lock.Lock()
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add1()
	wg.Wait()
	fmt.Println(x)
}
