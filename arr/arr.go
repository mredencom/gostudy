package main

import "fmt"

func main() {
	//数组定义明显的标志:1,有长度; 2,数组有具体的值
	var ipv4 = [4]uint8{192, 168, 0, 1}
	var ipv4em = [4]uint8{}
	//利用自定义一个数组 写出两种循环模式
	for k, v := range ipv4 {
		fmt.Println(k, ":", v)
	}
	fmt.Println("----------------------------------------")
	for i := 0; i < len(ipv4); i++ {
		fmt.Println(ipv4[i])
	}
	fmt.Println(ipv4)
	//空数组
	fmt.Println(ipv4em)
}
