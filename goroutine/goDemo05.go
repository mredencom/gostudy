package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

func workers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

//设置共享资源需要增加原始操作
var totals uint64
func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&totals, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()

	//fmt.Println(total.value)
	fmt.Println(totals)
}

