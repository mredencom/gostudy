package main

/**
练习goroutine   这里使用的ants包
*/
import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	//fmt.Println(i)
	n := i.(int32)
	//n := i.(string)
	atomic.AddInt32(&sum, n)
	//fmt.Printf("run with %d\n", n)
}
func main() {
	t1 := time.Now().Unix()
	runTimes := 10000000

	// 使用普通池
	var wg sync.WaitGroup
	//使用协程池
	p, _ := ants.NewPoolWithFunc(8, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	//最后销毁协程池
	defer p.Release()
	//循环提交任务
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	t2 := time.Now().Unix()
	fmt.Printf("运行协程: %d\n", p.Running())
	fmt.Printf("结果 %d\n", sum)
	fmt.Printf("使用时间 %d\n", t2-t1)
}
