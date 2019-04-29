//runner包管理处理定时任务的运行和声明周期
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner 在给定的超时时间内执行一组任务
//并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//interrupt 通道报告操作系统 发送的信号
	interrupt chan os.Signal

	//complete通道报告处理任务已经完成
	complete chan error

	//timeout报告处理任务超时
	timeout <-chan time.Time

	//task持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

//ErrTimeOut 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

//ErrInterrupt 会在接受到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

//New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将一个任务附加到 Runner 上。这个任务是一个
// 接收一个 int 类型的 ID 作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//start 执行所有任务, 监视通道的事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	//我这里使用goroutine执行任务
	go func() {
		//运行任务并告知complete情况
		r.complete <- r.run()
	}()
	//检查complete发出的信号
	select {
	//任务处理
	case err := <-r.complete:
		//返回完成的消息
		return err
	//任务是否超时
	case <-r.timeout:
		//返回超时的信息
		return ErrTimeout
	}
}

//执行已经存在的任务
func (r *Runner) run() error {
	//r.tasks 这个是一个方法slice
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			//返回中断信号的信息
			return ErrInterrupt
		}
		//执行任务
		task(id)
	}
	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//检查是否有中断信号的
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	//反之就可以继续运行了
	default:
		return false
	}
}
