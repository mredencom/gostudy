package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
	互斥锁的使用,可以解决竞争状态
		对 counterM 变量的操作在 Lock()和 Unlock()函数调用定义的临界
	区里被保护起来。使用大括号只是为了让临界区看起来更清晰，并不是必需的。同一时刻只有一
	个 goroutine 可以进入临界区。之后，直到调用 Unlock()函数之后，其他 goroutine 才能进入临
	界区。强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运
	行。当程序结束时，我们得到正确的值 4，竞争状态不再存在
*/
var (
	// counter 是所有 goroutine 都要增加其值的变量
	counterM int64
	// wg 用来等待程序结束 wgm sync.WaitGroup
	wgm sync.WaitGroup
	//mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wgm.Add(2)

	// 创建两个 goroutine
	go incCounterM()
	go incCounterM()

	// 等待 goroutine 结束
	wgm.Wait()
	fmt.Println("最终结果:", counterM)
}

func incCounterM() {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wgm.Done()
	for count := 0; count < 2; count++ {
		//加上锁时候同一时刻允许一个goroutine进入这个临界区
		mutex.Lock()
		//{
			// 捕获 counter 的值
			value := counterM
			//atomic.AddInt64(&counter,1)
			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()
			// 增加本地 value 变量的值
			value++
			// 将该值保存回 counter
			counterM = value
		//}
		mutex.Unlock()
		//释放锁,允许其正在等待的goroutine进入这个临界区
	}
}
