package main

import "fmt"

/**
	介绍new和make的区别
	总结：1.new关键字他不会默认初始化内存(置零)。2.而make会分配内存，会初始化.用于slice channel map
 */
func main() {
	p := new([]int) //p = nil
	fmt.Println(p)
	v := make([]int, 10, 50) //slice len 10 cap 50
	fmt.Println(v)
	(*p)[0] = 10 //这个会发生panic
	v[1] = 5     //这个正常
}
