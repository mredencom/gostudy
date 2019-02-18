package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	//输出内核启动进程ID 也就是顶级进程ID
	fmt.Println(os.Getppid())
	fmt.Println(os.Getgid())
	os.Exit(0)
	fmt.Printf("\n")
	//进程ID
	fmt.Println(os.Getpid())
	//杀死进程
	//signal := os.Kill
	//fmt.Println(signal)
	//创建一个管道的命令
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")
	if err := cmd0.Start(); err != nil {
		fmt.Println("发生错误")
		return
	}
	stdout0, err := cmd0.StderrPipe()
	if err != nil {
		fmt.Printf("你输入内容:%s\n", err)
		return
	}
	output0 := make([]byte, 30)
	n, err := stdout0.Read(output0)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", output0[:n])

	var outputBuf0 bytes.Buffer
	for {
		tempOutput := make([]byte, 5)
		n, err := stdout0.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("没有读到管道数据:%s\n", err)
				return
			}
		}
		if n > 0 {
			outputBuf0.Write(tempOutput[:n])
		}
	}
	fmt.Printf("%s\n", outputBuf0.String())


}
