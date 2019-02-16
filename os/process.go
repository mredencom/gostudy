package main

import (
	"fmt"
	"os"
)

func main() {
	//输出内核启动进程ID 也就是顶级进程ID
	fmt.Println(os.Getppid())
	fmt.Println("\n")
	//进程ID
	fmt.Println(os.Getpid())
	signal := os.Kill
	fmt.Println(signal)
}
