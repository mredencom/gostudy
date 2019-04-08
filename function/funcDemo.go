package main

import "fmt"

/**
	1.当一个函数有接受者这个函数就被称为"方法"
    2.Go 语言里有两种类型的接收者： "值接收者"和"指针接收者"
	3.这个接受者有类似于java和php中的类, 加入这个接受者就相当于这个类的方法.
*/
//方法能给用户定义的类型添加新的行为。方法实际上也是函数，只是在声明时，在关键字
//func 和方法名之间增加了一个参数
//user 定义一个用户类型
type user struct {
	name  string
	email string
}

//notify使用接受者实现了一个方法
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

//changeEmail使用指针接受者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

//main程序入口
func main() {
	//user 类型的值可以用来调用
	//使用值接受者声明的方法
	bill := user{"King", "123@163.com"}
	bill.notify()
	//指向user类型值的指针也可以用来调用
	//使用值接受者声明的方法
	lisa := &user{"Lisa", "Lisa@163.com"}
	lisa.notify()
	fmt.Println("------------")
	(*lisa).notify()
	fmt.Println("------------")
	//user类型的值可以用来调用
	//使用指针接受者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()
	(&bill).notify()
	//指向user类型值的指针可以用来调用
	//使用指针接受者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

}
