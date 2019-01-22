package main

import "fmt"

//map 的操作 实际就是hash table 官方给个称为map
func main() {
	//定义一个map 类型
	var ipSwitch = map[string]bool{}
	fmt.Println(ipSwitch)
	//向map插入数据 个人感觉向数组的键值对(json)
	ipSwitch["192.168.0.1"] = false
	ipSwitch["192.168.0.1"] = true
	ipSwitch["192.168.0.3"] = false
	//map 的键值对是唯一的
	fmt.Println(ipSwitch)
	//map 的键值的删除操作
	delete(ipSwitch, "192.168.0.1") //第一个参数是map 第二个参数是键
	fmt.Println(ipSwitch)
	ipSwitch["192.168.0.3"] = false
	ipSwitch["192.168.0.4"] = true
	ipSwitch["192.168.0.5"] = false
	ipSwitch["192.168.0.6"] = true
	ipSwitch["192.168.0.7"] = false
	fmt.Println("在循环删除之前:", ipSwitch)
	for k := range ipSwitch {
		delete(ipSwitch, k)
	}
	fmt.Println("在循环删除之后:", ipSwitch)

}
