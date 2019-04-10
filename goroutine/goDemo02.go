package main

import (
	"fmt"
	"runtime"
	"sync"
)

//创建wg 用来等待程序执行完成
var wg sync.WaitGroup

func main() {
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//这里添加两个计数器, 表示等待两个goroutine 完成
	wg.Add(2)
	//创建两个goroutine
	fmt.Println("Create Goroutine")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting to Finish")
	wg.Wait()
}

//printPrime 显示5000以内的素数
func printPrime(prefix string) {
	//在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
