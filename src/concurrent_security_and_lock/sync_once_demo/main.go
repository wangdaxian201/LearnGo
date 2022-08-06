package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover inner panic:%v\n", r)
		}
	}()
	for {
		x, ok := <-ch1
		if ok {
			ch2 <- x * x
		}
		if !ok {
			once.Do(func() {
				close(ch2)
			})
			break
		}

		//select {
		//case ch2 <- x * x:
		//	continue
		//default:
		//	break
		//}
	}

}

func main() {
	// 有时会出现一个panic, send on closed channel
	// 重启几次就会好
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover outer panic:%v\n", r)
		}
	}()
	a := make(chan int, 100)
	b := make(chan int, 100)
	//ctx, cancel := context.WithCancel(context.TODO())
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for res := range b {
		fmt.Println(res)
	}
	fmt.Println("exit")
	//time.Sleep(3 * time.Second)
}
