package main

import "fmt"

/**
*和&的区别 :
	& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
*/
type Rect struct {
	width  float64
	height float64
}

//取出面积值
func (r *Rect) size() float64 {
	return r.height * r.width
}
func main() {
	//1.&是取地址符号, 取到Rect类型对象的地址
	fmt.Println(&Rect{100, 100})
	//2.*可以表示一个变量是指针类型(r是一个指针变量)
	var r *Rect = &Rect{100, 100}
	fmt.Println(r)
	//3.*也可以表示指针类型变量所指向的存储单元 ,也就是这个地址所指向的值
	var rs *Rect = &Rect{100, 100}
	fmt.Println(&rs)
}
