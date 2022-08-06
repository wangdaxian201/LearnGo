package main

import (
	"fmt"
	"runtime"
	"sync"
)

// state sync waitGroup
// recommended sync.waitGroup
var wg sync.WaitGroup

func hello1(i int) {
	fmt.Println("hello 1, ", i)
	wg.Done()
}

func hello2() {
	fmt.Println("hello 2")
	wg.Done()
}

func main() {
	// start two goroutine
	// The goroutine assigned to multiple cores
	runtime.GOMAXPROCS(6)
	count := 10000
	wg.Add(count * 2)
	for i := 0; i < count; i++ {

		go func(i int) {
			fmt.Println("hello, ", i)
			wg.Done()
		}(i)
	}

	for i := 0; i < count; i++ {

		go func(i int) {
			fmt.Println("hello222, ", i)
			wg.Done()
		}(i)
	}
	fmt.Println("start learn concurrent...")
	fmt.Println("concurrent: At the same time perform multiple tasks!")
	// not recommended time.Sleep
	//time.Sleep(time.Second)
	wg.Wait()
}
