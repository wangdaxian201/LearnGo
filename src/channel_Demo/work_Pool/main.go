package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(Id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker: %d, start: %d \n", Id, job)
		results <- job * 2
		fmt.Printf("worker: %d, end: %d \n", Id, job)
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	wg.Add(3)

	// create 3 goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	// create five job write to jobs
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

}
