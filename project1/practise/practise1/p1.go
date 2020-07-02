package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		fmt.Fprintln(conn, "message")
		time.Sleep(5 * time.Second)
	}
	conn.Close()
}
