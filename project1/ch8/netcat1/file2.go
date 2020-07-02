package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000") // 连接TCP服务器
	if err != nil {
		log.Fatal(err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn) // 从conn到stdout
	mustCopy(conn, os.Stdin)     // 从stdin到conn ,主goroutine

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
