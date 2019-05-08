//work 包管理一个goroutine 池来完成工作
package work

import "sync"

//Worker 必须满足接口类型
//才能使用工作池
type Worker interface {
	Task()
}

//Pool 提供一个goroutine 池, 这个池可以完成
//任何提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			//轮询管道的work值 阻塞
			for w := range p.work {
				//然后执行管道里面的任务
				w.Task()
			}
			p.wg.Done()
		}()
	}
	//返回池的数据
	return &p
}

//Run 把任务提交到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//shutdown 等待所有goroutine 停止工作
func (p *Pool) Shutdown() {
	//关闭管道
	close(p.work)
	//等待任务完成
	p.wg.Wait()
}
