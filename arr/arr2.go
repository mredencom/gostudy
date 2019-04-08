package main

import "fmt"

func main() {
	var array = [2]*int{new(int), new(int)}
	*array[0] = 10
	*array[1] = 20
	fmt.Println(&array)
	var array1 = [5]string{}
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//这样可以吧array2 的值复制到array1 中
	array1 = array2
	fmt.Println(array1)
	//建立指针数组 声明一个三个元素都指针数组
	var array3 [3]*string

	array4 := [3]*string{new(string), new(string), new(string)}
	*array4[0] = "red"
	*array4[1] = "blue"
	*array4[2] = "green"

	array3 = array4
	fmt.Println(array3)

}
