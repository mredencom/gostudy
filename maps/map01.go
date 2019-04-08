package main

import "fmt"

func main() {
	//使用make创建一个映射(hash)
	dict := make(map[string]int)
	dicts := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict, dicts)
}
