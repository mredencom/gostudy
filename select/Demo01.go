package main

import (
	"fmt"
	"math/rand"
)

//测试goroutine 中select
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for {
			select {
			case x := <-c1:
				fmt.Println("从 c1 接受数据；", x)
			case c2 <- 100:
				fmt.Println("向 c2 发送数据")
			default:
				fmt.Println("c1 和 c2 都没什么可操作的")
			}
		}
	}()

	for i := 0; i < 500; i++ {
		rd := rand.Intn(2)
		switch rd {
		case 0:
			c1 <- 200
		case 1:
			<-c2
		}
	}
}
