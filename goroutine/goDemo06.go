package main

import (
	"fmt"
	"sync"
)
// 并发版本的hello world 四种书写方式
//https://www.goroutine.me/2019/08/27/goroutine.html
func main() {
	g04()
}

//第一个版本
func g01() {
	done := make(chan int)
	go func() {
		fmt.Println("你好, 世界")
		<-done
	}()
	done <- 1
}

//第二个版本
func g02() {
	done := make(chan int, 1) // 带缓存的管道
	go func() {
		fmt.Println("你好, 世界")
		done <- 1
	}()
	<-done
}

//第三个版本
func g03() {
	done := make(chan int, 10)
	for i := 1; i < cap(done); i++ {
		go func() {
			fmt.Println("Hello World %", i)
			done <- i
		}()
	}
	for i := 1; i < cap(done); i++ {
		fmt.Println(<-done, "done")
	}
}

//第四个版本
func g04() {
	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello World", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
