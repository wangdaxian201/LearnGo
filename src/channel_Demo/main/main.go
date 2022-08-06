package main

import "fmt"

func write(c chan int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func read(c1, c2 chan int) {
	for {
		x, ok := <-c1
		if !ok {
			break
		}
		c2 <- x * x
	}

	close(c2)
}

func main() {
	fmt.Println("learn channel....")
	fmt.Println("channel operation list: (send)、(receive) (close)")
	fmt.Println("channel need make initialize")

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go write(ch1)
	go read(ch1, ch2)

	for ret := range ch2 {
		fmt.Println("ch2: ", ret)
	}

}
