package main

import "fmt"

//就是结构体可以嵌套使用
//结构体的嵌套 实现接口方法, 如果内部类型没有实现就会调用外部实现方法(内部提升到外部),如果内部实现了就不会调用外部的方法(内部不会提升到外部). (这个有点像继承机制)
/**
	Go 语言允许用户扩展或者修改已有类型的行为。这个功能对代码复用很重要，在修改已有
类型以符合新类型的时候也很重要。这个功能是通过嵌入类型（type embedding）完成的。嵌入类
型是将已有的类型直接声明在新的结构类型里。被嵌入的类型被称为新的外部类型的内部类型。
通过嵌入类型，与内部类型相关的标识符会提升到外部类型上。这些被提升的标识符就像直
接声明在外部类型里的标识符一样，也是外部类型的一部分。这样外部类型就组合了内部类型包
含的所有属性，并且可以添加新的字段和方法。外部类型也可以通过声明与内部类型标识符同名
的标识符来覆盖内部标识符的字段或者方法。这就是扩展或者修改已有类型的方法
*/
// notifier 是一个定义了
// 通知类行为的接口
type notifiers interface {
	sendNotify()
}

//定义一个用户的结构体
type users struct {
	name  string
	email string
}

// notify 实现了一个可以通过 user 类型值的指针
// 调用的方法
func (u *users) sendNotify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//admins 代表一个拥有权限管理员用户
type admins struct {
	users //外部类型
	level string//内部类型
}

func (u *admins) sendNotify() {
	fmt.Printf("Sending admin email to %s<%s>%s\n",
		u.name,
		u.email,
		u.level)
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotifications(s notifiers) {
	s.sendNotify()
}
func main() {
	//创建一个 admin 用户
	ad := admins{users: users{
		name:  "King Smith",
		email: "smith@yahoo.com",
	},
		level: "superman",
	}
	//我们可以直接访问内部类型的方法
	ad.users.sendNotify()
	//内部类型的方法也被提升到外部类型
	ad.sendNotify()
	// 给 admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到
	// 外部类型
	sendNotifications(&ad)
}
