package main

import (
	"fmt"
	"sync"
	"time"
)

/**
我们现在玩一个跑步接力赛游戏,
*/
//创建一个wg用来等待程序的结束
var wg1 sync.WaitGroup

func main() {
	//创建一个无缓冲的通道
	baton := make(chan int)
	//为最后一个跑步者将计数加1
	wg1.Add(1)
	//为第一个跑者开一个协程 开始运动
	go Runner(baton)
	//开始比赛
	baton <- 1
	//等待程序结束
	wg1.Wait()

}

//模拟选手跑步
func Runner(baton chan int) {
	var newRunner int
	//等待接力棒
	runner := <-baton
	fmt.Printf("选手 %d 带着接力棒开始跑\n", runner)
	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("向选手 %d 传递接力棒\n", newRunner)
		//选手开始接力赛
		go Runner(baton)
	}
	time.Sleep(100 * time.Microsecond)
	if runner == 4 {
		fmt.Printf("选手 %d 完成比赛\n", runner)
		wg1.Done()
		return
	}
	// 将接力棒交给下一位跑步者
	fmt.Printf("选手 %d 传递接力棒给选手 %d\n",
		runner,
		newRunner)
	//新的一位选手带着接力棒继续前进
	baton <- newRunner
}
