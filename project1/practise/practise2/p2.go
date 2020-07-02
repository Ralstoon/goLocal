package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		str := scanner.Text()
		os.Stdout.WriteString(str)
	}
	// readBytes := make([]byte, 1)
	// var buffer bytes.Buffer
	// for {
	// 	_, err := conn.Read(readBytes)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("Client EOF err!")
	// 		} else {
	// 			fmt.Println("Read Err!")
	// 		}
	// 		break
	// 	}
	// 	readByte := readBytes[0]
	// 	buffer.WriteByte(readByte)
	// 	if readByte == '\n' {
	// 		os.Stdout.Write(buffer.Bytes())
	// 		buffer.Reset()
	// 	}
	// }

}
