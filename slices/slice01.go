package main

import "fmt"

func main() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println(cap(source))
	slice := source[1:3:5]
	slice = append(slice, "sss")
	fmt.Println(source, slice)
}
