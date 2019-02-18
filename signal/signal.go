package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/**
	进程之间的通信方式: 信号
 */
func main() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGQUIT, syscall.SIGINT}
	fmt.Printf("set notification for %s...[sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Receive a signal from sigRecv1:%s\n", sig)
		}
		fmt.Printf("end.[sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Receive a signal from sigRecv2:%s\n", sig)
		}
		fmt.Printf("end.[sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("等待两秒...")
	time.Sleep(2 * time.Second)
	fmt.Printf("stop notification")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("done.[sigRecv1]\n")
	wg.Wait()
}

func sigRecv1() {
	sigRecv := make(chan os.Signal, 1)
	//加入两个信号值 在linux 下可以使用命令 kill -l 查看
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	//使用notify发送信号
	signal.Notify(sigRecv, sigs...)
	for sig := range sigRecv {
		//输出
		fmt.Printf("Received a signal: %s\n", sig)
	}
}
func sigRecv11() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGQUIT, syscall.SIGINT}
	fmt.Printf("set notification for %s...[sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)
}

func sigRecv22() {
	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

}
