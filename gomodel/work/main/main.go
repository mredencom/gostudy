//这个实例展示如何使用work包
//创建一个goroutine 池并完成工作
package main

import (
	"gostudy/gomodel/work"
	"log"
	"sync"
	"time"
)

//建立一组将显示的名称
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}
//namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

//Task 实现work接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//使用两个goroutine创建工作池
	p := work.New(5)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			//创建一个namePrinter并提供指定名字
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	//等待
	wg.Wait()
	//关闭
	p.Shutdown()
}
