package main

import (
	//"code-master/chapter7/patterns/runner"
	"gostudy/gomodel/runner"
	"log"
	"os"
	"time"
)

//定义timeout3秒必须处理完成任务
const timeout = 3 * time.Second

//主程序入口
func main() {
	log.Println("工作开始...")
	//分配超时时间
	r := runner.New(timeout)
	//runner.New()
	//加入要执行的任务
	r.Add(createTask(), createTask(), createTask())
	//这里开启任务和知道任务的
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("任务超时")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("任务退出")
			os.Exit(2)
		}
	}
	log.Println("工作结束...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("任务ID #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
