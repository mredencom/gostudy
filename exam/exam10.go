package main

import (
	"fmt"
	"reflect"
)

type People1 interface {
	Show()
}
type Student struct {
}

func (stu *Student) Show() {

}
func live() People1 {
	var stu *Student
	return stu
}
func main() {
	fmt.Println(live())
	fmt.Println(reflect.TypeOf(live()))
	if live() == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
