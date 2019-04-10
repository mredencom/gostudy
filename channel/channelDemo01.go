package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//wg用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)
	// 启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)
	// 发球
	court <- 1
	// 等待游戏结束35
	wg.Wait()
}

//模拟选手打球传球
func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了 通道没数据
			fmt.Printf("Player %s Won\n", name)
			return
		}
		//选随机数 然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}
		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}
