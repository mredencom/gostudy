package main

import "fmt"
/**
	使用go vet 检查代码的一些静态错误信息
	go vet 工具不能让开发者避免严重的逻辑错误，或者避免编写充满小错的代码。不过，正
	像刚才的实例中展示的那样，这个工具可以很好地捕获一部分常见错误。每次对代码先执行 go
	vet 再将其签入源代码库是一个很好的习惯
	http://note.youdao.com/noteshare?id=1f3294b2c07ba45f2190130653b73738&sub=wcp1553574279096966 可以查看这里使用
*/
func main() {
	fmt.Printf("The quick brown fox jumped over lazy dogs %s", 3.14)
}
