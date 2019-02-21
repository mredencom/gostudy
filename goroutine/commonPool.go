package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

func demoFunc() {
	time.Sleep(10 * time.Microsecond)
	fmt.Println("你好...")
}
func main() {
	defer ants.Release()
	runTimes := 10000000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i <= runTimes; i++ {
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	fmt.Printf("运行中协程: %d\n", ants.Running())
	fmt.Printf("运行完成!\n")
}
