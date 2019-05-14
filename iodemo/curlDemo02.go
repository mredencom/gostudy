package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

//这个例子程序展示来自不同标准库不同函数是如何使用io.writer 和 io.Reader 接口
//写一个简单版本curl
func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	//创建文件来保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	//最后关闭file
	defer file.Close();
	//使用MultiWriter 这样就可以同时向文件和标准输出设备
	//进行写入操作
	dest := io.MultiWriter(os.Stdout, file)
	//读取响应内容,并吸入两个目的地
	_, _ = io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}

}
