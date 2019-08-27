package main

import "fmt"

//https://www.goroutine.me/2019/08/27/goroutine.html
func main() {
	//done := make(chan int,1)
	done := make(chan int)
	go func() {
		fmt.Println("Hello World")
		//done <- 1
		<-done
	}()
	//<-done
	done <- 1
}
