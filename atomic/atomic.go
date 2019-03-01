package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	计数器使用
	互斥量 锁
 */

type atomicInt struct {
	value int
	lock  sync.Mutex
}

//计数器 原子性操作 属于线程安全的.
func (a *atomicInt) increment() {
	fmt.Println("安全增加")
	for i := 1; i <= 10000; i++ {
		func() {
			a.lock.Lock()
			defer a.lock.Unlock()
			a.value++
		}()
	}
}

//取数据
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atomicInt
	//a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(2 * time.Microsecond)
	fmt.Println(a.get())
}
