package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("客户端连接错误❌ ! " + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " : Client connected!")

	onMessageRecived(conn)
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	_, _ = conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("读取字符串")
		fmt.Println(msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		//time.Sleep(time.Second * 2)

		fmt.Println("写入...")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
