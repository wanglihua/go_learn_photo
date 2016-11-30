//tcp server
package main

import (
	"fmt"
	"net"
)

const (
	ip            = ""
	terminal_port = 3333
	web_port      = 4444
)

func main() {
	go start_listen_terminal()
	go start_listen_web()

	var input []byte
	fmt.Scan(&input)
}

var terminal_conn_map map[string]*net.TCPConn = make(map[string]*net.TCPConn)

func start_listen_terminal() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), terminal_port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待 terminal 客户端连接...")
	start_terminal_server(listen)
}

func start_terminal_server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受 terminal 客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("terminal 客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				if err != nil {
					fmt.Println("读取 terminal 客户端数据错误:", err.Error())
					break
				}
				fmt.Println("terminal 客户端发来数据:", string(data[0:i]))
				if len(data) > 5 && data[0] == 't' && data[1] == 'e' && data[2] == 'r' && data[3] == 'm' {
					if data[4] == '1' {
						terminal_conn_map["term1"] = conn
					}
					if data[4] == '2' {
						terminal_conn_map["term2"] = conn
					}
				}

				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}
		}()
	}
}

func start_listen_web() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), web_port, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待 web 客户端连接...")
	start_web_server(listen)
}

func start_web_server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受 web 客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("web 客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Println("web 客户端发来数据:", string(data[0:i]))
				if string(data[0:i]) == "1" && terminal_conn_map["term1"] != nil {
					terminal_conn_map["term1"].Write([]byte{'t', 'e', 'r', 'm', '1'})
				}
				if string(data[0:i]) == "2" && terminal_conn_map["term2"] != nil {
					terminal_conn_map["term2"].Write([]byte{'t', 'e', 'r', 'm', '2'})
				}
				if err != nil {
					fmt.Println("读取 web 客户端数据错误:", err.Error())
					break
				}
				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}

		}()
	}
}
