package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int64
	// wg 用来等待程序结束16 wg sync.WaitGroup
	wgs sync.WaitGroup
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wgs.Add(2)

	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wgs.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wgs.Done()
	for count := 0; count < 2; count++ {
		// 捕获 counter 的值
		//value := counter

		atomic.AddInt64(&counter,1)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
		// 增加本地 value 变量的值
		//value++

		// 将该值保存回 counter
		//counter = value
	}
}
