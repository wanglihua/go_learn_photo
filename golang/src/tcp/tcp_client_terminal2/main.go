//tcp client terminal 2
package main

import (
	"fmt"
	"net"
	"time"
)

const (
	addr = "127.0.0.1:3333"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	Client(conn)
}

func Client(conn net.Conn) {
	//sms := make([]byte, 128)
	for {
		time.Sleep(1 * time.Second)
		//fmt.Print("请输入要发送的消息:")
		//_, err := fmt.Scan(&sms)
		//if err != nil {
		//	fmt.Println("数据输入异常:", err.Error())
		//}
		conn.Write([]byte{'t', 'e', 'r', 'm', '2'})
		//conn.Write(sms)
		conn.Write([]byte(" Hello from term2"))
		buf := make([]byte, 128)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println(string(buf[0:c]))
	}

}
