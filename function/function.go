package main

import (
	"errors"
	"fmt"
)

func main() {
	//result, err := div(60, 2)
	//if err != nil {
	//
	//} else {
	//	fmt.Println(result)
	//}
	//

}

type myInt int //定义类型

func (i myInt) add(another int) myInt {
	i = i + myInt(another)
	return i
}

//defer 捕获异常
func errs() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	divide()
	fmt.Printf("这里应该执行不到")
}
func divide() {
	var x = 30
	var y = 0
	var c = x / y
	fmt.Println(c)

}

func div(dividend int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("除数不能为0")
		return
	}
	result = dividend / divisor
	return result, nil
}

//用于定义一个二元操作类型
type binaryOperation func(operand1 int, operand2 int) (result int, err error)

//用于定义执行二元操作
func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("invalid binary operation function")
		return
	}
	return bop(op1, op2)
}
