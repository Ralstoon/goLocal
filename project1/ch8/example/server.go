package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	// conn.LocalAddr()
	ch := make(chan string)
	defer conn.Close()
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println("no reaction,Closed!")
				conn.Close()
				return
			case str := <-ch:
				echo(conn, str, 1*time.Second)
			}
		}
	}()
	input := bufio.NewReader(conn)
	for {
		readBytes, err := input.ReadBytes(' ')
		if err != nil {
			fmt.Println(err)
		}
		str := string(readBytes)
		fmt.Println(str)
		ch <- str
	}
	// for input.Scan() {
	// 	str := input.Text()
	// 	if str == "exit" {
	// 		break
	// 	}

	// }
}

func echo(c net.Conn, input string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(input))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", input)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(input), '&')
}
