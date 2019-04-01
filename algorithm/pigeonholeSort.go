package main

import "fmt"

/**
鸽巢排序是桶排序的一种，反正是变着花样玩，顾名思义，就是一排鸽巢，看里面有几个鸽巢，然后遍历这些鸽巢，打印出来就好，排序之前得先知道区间和最大值

我们通过一个简单的例子来讲解，讲解完之后希望大家可以对鸽巢排序有一个深入的了解

比如有数组a=[2,7,5,9,8,8]，我们需要对这个数组进行排序

嗯，看到数组之后，我们发现这是一个最大值不超过10的数组，那么我们定区间为0-10，定义一个下标0-10这样一个11位数组b,初始化值为0

然后遍历已知数组a，通过a的一项一项的值和我们定义的数组b的下标进行对应

遍历开始

取a[0]=2，然后操作b[2]=1

取a[1]=7，然后操作b[7]=1

取a[2]=5，然后操作b[5]=1

取a[3]=9，然后操作b[9]=1

取a[4]=8，然后操作b[8]=1

取a[5]=8，然后操作b[8]=2（注意了哦）

这样就对号入座了，之后遍历b，如果b中某个小标对应的值是多个，则遍历多次，把不是0的给打印出来，结果就是我们想要的了
*/
//不稳定排序
func pigeonholeSort(theArray []int) {
	vl := 0
	//取出最大值
	for _, v := range theArray {
		if v > vl {
			vl = v
		}
	}
	//由于go中不能使用变量作为数组定义长度,所以使用切片
	slice := make([]int, vl+1)
	for _, v := range theArray {
		slice[v]++
	}
	for i := range slice {
		if slice[i] != 0 {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	fmt.Println(vl)
}
func main() {
	var theArray = []int{1, 0, 2, 10, 9, 70, 5, 6, 3}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	pigeonholeSort(theArray)
}
