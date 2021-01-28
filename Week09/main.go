package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		return;
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("err: %s\n", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	msg := make(chan string, 1)
	defer conn.Close()
	defer close(msg)

	go send_msg(conn, msg)
	get_msg(conn, msg)
}

func get_msg(conn net.Conn, message chan string) {
	reader := bufio.NewScanner(conn)
	for reader.Scan() { // 当对端断连，这里会退出
		content := reader.Text()
		// 忽略掉消息处理
		log.Printf("Receive message: %s\n", content)
		message <- content
	}
}

func send_msg(conn net.Conn, msg <-chan string) {
	for content := range msg {
		fmt.Println(conn, content)
	}
}
