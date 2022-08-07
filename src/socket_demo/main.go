package main

import (
	"fmt"
	"net"
	"socket_demo/server"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8000")
	fmt.Println("listen : 0.0.0.0:8000....")
	if err != nil {
		fmt.Printf("listen faild, err %v\n", err)
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept faild, err %v\n", err)
			continue
		}
		go server.Process(accept)
	}

}
