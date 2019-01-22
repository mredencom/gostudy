package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//声明并初始化带缓冲的读取器
	//准备从标准输入读取输入内容
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("发现一个错误: %s\n", err)
	} else {
		input = input[:len(input)-1]
		fmt.Printf("你好: %s\n", input)
	}
}
