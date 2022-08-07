package server

import (
	"bufio"
	"fmt"
	"net"
)

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from client failed,  err %v\n", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("recv: ", recvStr)
		_, _ = conn.Write([]byte("ok"))
		if err != nil {
			return
		} // 返回给客户端信息
	}

}
