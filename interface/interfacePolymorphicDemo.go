package main

import "fmt"

/**
	接口中多态的实现
*/
//定义一个接口
type notifierMessage interface {
	notifyMessage()
}

//定义一个用户的结构
type userInfo struct {
	name  string
	email string
}

//notify 使用指针接收者实现了 notifier 接口
func (u *userInfo) notifyMessage() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//adminInfo 定义了程序员里的管理员
type adminInfo struct {
	name  string
	email string
}

//notify 使用指针接收者实现了 notifier 接口
func (a *adminInfo) notifyMessage() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

//并发送通知
func sendNotificationMessage(n notifierMessage) {
	n.notifyMessage()
}
func main() {
	// 创建一个 user 值并传给 sendNotification
	bills := userInfo{"Bill", "bill@email.com"}
	sendNotificationMessage(&bills)

	// 创建一个 admin 值并传给 sendNotification
	lisa := adminInfo{"Lisa", "lisa@email.com"}
	sendNotificationMessage(&lisa)
}
