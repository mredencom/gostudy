package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

const (
	URL = "127.0.0.1:12981"
)

type Ah int

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//乘法
func (t *Ah) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

//除法
func (t *Ah) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为零")
	}
	quo.Quo = args.A / args.B //除法
	quo.Rem = args.A % args.B //求余
	return nil
}
func main() {
	ah := new(Ah)
	//注册
	_ = rpc.Register(ah)
	rpc.HandleHTTP()
	err := http.ListenAndServe(URL, nil)
	if err != nil {
		fmt.Println("发生错误:", err.Error())
	}

}
