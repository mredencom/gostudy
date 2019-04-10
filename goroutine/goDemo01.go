//创建实例进程goroutine
//以及调度器的行为
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//main程序入口
func main() {
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)
	//wg用来等待程序完成
	//计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("协程开始")
	//声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()
	//声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()
	// 等待 goroutine 结束
	fmt.Println("等待完成")
	wg.Wait()

	fmt.Println("程序完成")
}
