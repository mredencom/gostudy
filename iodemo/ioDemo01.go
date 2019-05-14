package main

import (
	"bytes"
	"fmt"
	"os"
)

//这个例子程序展示来自不同标准库不同函数是如何使用io.writer接口
//main应用程序入口
func main() {
	//创建一个Buffer值,并将一个字符串写入Buffer
	//使用io.Writer和write
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	//使用Fprintf来将一个字符串拼接到Buffer里,将bytes.Buffer的地址作为io.writer类型值写入
	_, _ = fmt.Fprintf(&b, "World!")
	//将Buffer的内容输出到标准输出设备.将os.File值的地址作为io.Writer类型值传入
	_, _ = b.WriteTo(os.Stdout)
}
