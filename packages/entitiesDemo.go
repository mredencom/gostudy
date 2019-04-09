package main

import (
	"fmt"
	"gostudy/packages/entities"
)

func main() {
	//u := entities.User{Name: "Bill", Email: "bill@163.com"}
	//fmt.Printf("User: %v\n", u)
	//嵌入类型 即使嵌入类型未公开
	a := entities.Admin{Rights: 111}
	a.Name = "KING"
	a.Email = "email@qq.com"
	fmt.Printf("User: %v\n", a)
}
