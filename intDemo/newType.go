package main

/**
基于int64声明一个新的类型
*/
type Duration int64

func main() {
	//给不同类型的变量赋值会产生编译错误 注意
	//会报错:cannot use int64(1000) (type int64) as type Duration in assignment
	//var dur Duration
	//dur = int64(1000)

}
