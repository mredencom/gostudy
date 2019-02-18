package main

import (
	"runtime"
	"sync"
)

func sum(ch chan int ,wg *sync.WaitGroup) {
	//fmt.Println(arrays)
	//sum := 0
	//for _, array := range arrays {
	//	sum += array
	//}
	for i:=0;i<10000000 ;i++  {
		ch <- i
		wg.Done()
	}

}

func main() {
	//创建信号灯(可以理解成队列)
	var wg sync.WaitGroup
	arrayChan := make(chan int, 20)
	cpuNum := runtime.NumCPU()
	for t := 0; t < cpuNum; t++ {
		wg.Add(1)
		go sum(arrayChan,&wg)
	}
	<-arrayChan

	wg.Wait()
}
