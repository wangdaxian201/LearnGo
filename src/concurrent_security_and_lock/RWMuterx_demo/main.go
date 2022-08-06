package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.RWMutex
)

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

func read() {
	lock.RLock()
	time.Sleep(10 * time.Millisecond)
	lock.RUnlock()
	wg.Done()
}

func write() {
	lock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	lock.Unlock()
	wg.Done()
}
