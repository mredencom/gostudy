package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func runs(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		c <- i
	}
}

func main() {
	t1 := time.Now().Unix()
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(8) //设置cpu的核的数量，从而实现高并发
	c := make(chan int)
	for i := 0; i < 8; i++ {
		go runs(c, &wg)
	}
	var data []int
	for i := range c {
		data = append(data, i)
		wg.Done()
	}
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)
	wg.Wait()
}
