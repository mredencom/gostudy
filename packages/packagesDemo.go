package main

/**
	当一个标识符的名字以小写字母开头时，这个标识符就是未公开的，即包外的代码不可见。
如果一个标识符以大写字母开头，这个标识符就是公开的，即被包外的代码可见。让我们看一下
导入这个包的代码
*/
import (
	"fmt"
	"gostudy/packages/counters"
)

//这个示例程序展示无法从另一个包里
//访问未公开的标识符
//main程序入口
func main() {
	// 创建一个未公开的类型的变量
	// 并将其初始化为 10
	counter := counters.New(10)

	// ./listing64.go:15: 不能引用未公开的名字  首字母大写是包公开变量  小写是私有变量
	// counters.alertCounter
	// ./listing64.go:15: 未定义： counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
