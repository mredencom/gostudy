package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/**
建立一个创建执行任务的方法(取数据)
 */
func doWork(w worker) {
	var datas []int
	for n := range w.in {
		//fmt.Printf("数据: %d \n", n)
		datas = append(datas, n)
		w.done1()
	}
}

/**
	创建一个结构体
 */
type worker struct {
	in    chan int
	done1 func()
}

/**
	创建一个任务方法
 */
func createWorker(wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done1: func() {
			wg.Done()
		},
	}
	//开任务
	go doWork(w)
	return w
}

/**
	使用循环创建任务
 */
func chanDemo() {
	//创建信号灯(可以理解成队列)
	var wg sync.WaitGroup
	//slice
	var workers []worker
	//取cpu个数
	numCpu := runtime.NumCPU()
	//根据运算cpu个数创建任务
	for i := 0; i < numCpu; i++ {
		//workers[i] = createWorker(i, &wg)
		workers = append(workers, createWorker(&wg))
	}
	//os.Exit(0)
	//增加信号值
	//wg.Add(numCpu * 1000)
	for _, worker := range workers {
		//worker.in <- i
		//向每一个任务加入1000条数据
		for i := 0; i < 10000000; i++ {
			//增加信号值
			wg.Add(1)
			//向结构体加入数据
			//worker.in <- i
			worker.in <- rand.Intn(10000)
		}
		//createData(worker)
	}
	//等待任务执行完(阻塞)
	wg.Wait()
}

func main() {
	//执行任务
	t1 := time.Now().Unix()
	chanDemo()
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)
}
