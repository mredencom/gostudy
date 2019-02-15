package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
	简单的聊天机器人
 */
func main() {
	//准备冲标准输入读取数据源
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入您的名字:")
	//读取数据知道碰到\n
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("出现一个错误: %s\n", err)
		// 异常退出。
		os.Exit(1)
	} else {
		// 用切片操作删除最后的 \n 。
		name := input[:len(input)-1]
		fmt.Printf("您好, %s! 我有什么可以帮助您?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("出现一个错误: %s\n", err)
			continue
		}
		input = input[:len(input)-1]
		// 全部转换为小写。
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出。
			os.Exit(0)
		case "1":
			fmt.Println("你是想学习阿拉伯数字吗?")
		default:
			fmt.Println("对不起!")
		}
	}
}
