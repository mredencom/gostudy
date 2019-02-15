package main

import "fmt"

func main() {
	//outerFunc()
	//
	//var str = "你好的的12aaA江涛"
	//for _, v := range str {
	//	//fmt.Println(i,v)
	//	fmt.Printf("Unicode: %c  %d\n", v, v)
	//}
	printNumber()
}

/**
	测试defer
 */
func outerFunc() {
	defer fmt.Println("函数执行结束前一刻会打印")
	fmt.Println("第一个会被打印")
}

/*
	测试defer 2
 */
func printNumber() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d", n)
		}(i * 2)
	}
}
