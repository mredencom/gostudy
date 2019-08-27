package main

import "fmt"

var p = f()

func main() {
	//fmt.Println(f() == f()) //false
	v := 1
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
}

func f() *int {
	v := 1
	return &v
}

func incr(c *int) int {
	*c++ //非常注意：这里的增加c指向的值，不会改变c指针
	return *c
}
func init() {

}
func init() {


}