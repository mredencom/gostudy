package main

import "fmt"

/**
 *切片(slice)基本信息 操作
 */
func main() {
	//定义一个slice
	//slice特征: 1,不需要定义长度; 2,slice内部数据类型相同 3,空的slice内部没有值是nil
	var ips = []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
	//一个切片类型的零值总是为nil 此零值的长度和容量都为0
	var ipsnil = []string{}
	fmt.Println(ipsnil)
	fmt.Println(cap(ips), len(ips))
	//打印ips slice的值
	fmt.Println(ips)
	//向ips slice 追加值
	ips = append(ips, "192.168.0.4", "192.168.0.5")
	fmt.Println(ips)

}
