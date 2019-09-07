package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Chmod("new path",0666)
	//log.Fatal(err.(syscall.Errno))
	//fmt.Println(err.(syscall.Errno))
	fmt.Println(err)
}
