package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp/client/main.go

// 客户端
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入你要发送给服务端的信息 (Q退出): ")
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("收到服务端的response", string(buf[:n]))
	}
}
