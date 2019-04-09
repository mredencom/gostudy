package main

import "fmt"

/**
方法集(一般在接口使用) (限制:所以值的方法集只包括了使用值接收者实现的方法)
	1.规范里描述的方法集
		Values              Methods Receivers
		-----------------------------------------------
		T                       (t T)
		*T                      (t T) and (t *T)
			展示了规范里对方法集的描述。描述中说到， T 类型的值的方法集只包含值
		接收者声明的方法。而指向 T 类型的指针的方法集既包含值接收者声明的方法，也包含指针接收
		者声明的方法。从值的角度看这些规则，会显得很复杂。让我们从接收者的角度来看一下这些规
		则，
	2.从接收者类型的角度来看方法集
		Methods                 Receivers Values
		-----------------------------------------------
		(t T)                       T and *T
		(t *T)                      *T
			展示了同样的规则，只不过换成了接收者的视角。这个规则说，如果使用指
		针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。如果使用值
		接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。就能理解出现编译错误的原因了，
*/
//notifier定义了一个通知接口

type notifier interface {
	notify()
}

//user 定义了一个用户类型
type user struct {
	name  string
	email string
}

//notify 是使用指针接受者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}
func sendNotification(n notifier) {
	n.notify()
}
func main() {
	//创建一个用户类型的值,并发送通知
	u := user{"Bill", "Bill@email.com"}
	//u 类型是user类型  但是 sendNotification 是notifier接口
	sendNotification(&u)
}
