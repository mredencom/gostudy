package main

import "fmt"

//以下实例通过 Go 语言的递归函数实现斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func main() {
	var i int
	for i = 0; i < 15; i++ {
		fmt.Printf("%d\t\n", fibonacci(i))
	}
}
