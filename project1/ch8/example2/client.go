package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func read(conn io.Reader, spliter byte) ([]byte, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return nil, err
		}
		readByte := readBytes[0]
		if readByte == spliter {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.Bytes(), nil
}

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	toClose := make(chan bool)
	// go func(conn net.Conn) {
	// 	for {
	// 		// conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	// 		strReq, err := read(conn, '\n')
	// 		fmt.Println("line 40 :", string(strReq), err)
	// 		if err != nil {
	// 			if err == io.EOF {
	// 				fmt.Println("Server has closed the connection42")
	// 			} else {
	// 				fmt.Printf("Read Error44:%s\n", err)
	// 			}
	// 			toClose <- true
	// 			break
	// 		}
	// 		// toBroadcast <- strReq
	// 		os.Stdout.Write(strReq)
	// 	}
	// }(conn)
	go func(conn net.Conn) {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			str := scanner.Text()
			os.Stdout.WriteString(str)
		}
		toClose <- true
	}(conn)

	for {
		select {
		case <-toClose:
			fmt.Println("server has closed client conn! Finished")
			return
		default:
			for {
				// conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				strReq, err := read(os.Stdin, '\n')
				if err != nil {
					fmt.Println(err)
					break
				}
				// toBroadcast <- strReq
				fmt.Println("read from stdin:\t", string(strReq))
				conn.Write(strReq)
			}
		}
	}

	// for s := range toBroadcast {
	// 	os.Stdout.WriteString(s)
	// }

	// conn.SetWriteDeadline(time.Now().Add(time.Second * 5)) // Server端断开连接，这边做超时处理
	// time.Sleep(6 * time.Second)

	// go func() {
	// 	arr := make([]byte, 8)
	// 	for {
	// 		_, err := conn.Read(arr)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			conn.Close()
	// 			return
	// 		}
	// 	}
	// 	os.Stdout.Write(arr)
	// }()

	// go mustCopy(os.Stdout, conn)
	// go func() {
	// 	for {
	// 		_, err = conn.Read(make([]byte))
	// 		if err != nil {--
	// 			conn.Close()
	// 		}
	// 	}
	// }()

	// mustCopy(conn, os.Stdin)
}

func mustCopy(des io.Writer, src io.Reader) {
	if _, err := io.Copy(des, src); err != nil {
		log.Fatal(err)
	}
}
