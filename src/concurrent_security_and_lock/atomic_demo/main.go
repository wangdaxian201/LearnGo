package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x  int32
	l  sync.Mutex
	wg sync.WaitGroup
)

func Add() {
	// 普通add函数
	x++
	wg.Done()
}

func mutexAdd() {
	// 互斥锁操作
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt32(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	// code body
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		//go Add()  // 普通版add 函数, 未使用并发安全的操作
		//go mutexAdd() // 互斥锁版本add函数, 并发安全, 但是加锁开销大
		atomicAdd() // 原子操作版add函数 是并发安全的 性能优于加锁版

	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start))
}
