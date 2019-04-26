package main

import (
	"fmt"
	"net/rpc"
)

//监听rpc服务地址
const (
	URLS = "127.0.0.1:12981"
)

//构造参数
type Arg struct {
	A, B int
}

//求商和求余返回结构体
type Quotients struct {
	Quo, Rem int
}

func main() {
	//监听rpc地址
	client, err := rpc.DialHTTP("tcp", URLS)
	if err != nil {
		fmt.Println(err.Error())
	}
	i := Arg{10000, 3}
	q := new(Quotients)
	//var reply int
	//呼叫远程的方法 参数1 远程方法  参数2 使用的参数 第三个参数接受值,
	//err = client.Call("Ah.Multiply", &i, &reply)
	err = client.Call("Ah.Divide", &i, &q)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*q)
	}
}
