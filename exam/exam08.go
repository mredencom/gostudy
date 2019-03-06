package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出错了！", err)
		}
	}()
	t := UserAges{}
	go t.Add("小明", 12)
	fmt.Println("小明12岁")

	//go func() {
	//	t.Add("小明", 12)
	//	fmt.Println("小明12岁")
	//}()
	//time.Sleep(time.Microsecond)
	//t.Add("小明", 12)
	//t.Get("你")
}
