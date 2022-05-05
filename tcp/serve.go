package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func progress(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:]) //读取数据
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			return
		}
		receStr := string(buf[:n])
		fmt.Println("收到client端发送来的数据", receStr)
		conn.Write([]byte(receStr)) //发送数据
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed ,err；", err)
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,err：", err)
			continue
		}
		go progress(conn) //启动一个goroutine处理连接
	}
}
