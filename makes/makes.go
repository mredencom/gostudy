package main

import "fmt"

//make 关键字 用于定义切片 数组 字典(map) 通道类型(channel)
func main() {
	//使用make定义一个切片 第一个100元素长度为100的 第二个100预留100个元素的存储空间
	ipsArr := make([]int, 100, 100)
	//取地址
	//println(&ipsArr)
	fmt.Println(ipsArr)
	//定义一个通道  make(chan type, cap) cap容量
	//一个有缓存的通道
	ch := make(chan int, 10)
	//创建一个无缓存的通道
	ch1 := make(chan int)
	//返回通道地址
	fmt.Println(ch, "\n", ch1)

}
