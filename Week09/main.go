package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error:%v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan []byte, 10)
	go read(conn, ch)
	go write(conn, ch)
}

func read(conn net.Conn, ch chan<- []byte) {
	// 终止读取后，关闭通道
	defer close(ch)

	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Printf("read error:%v", err)
			return
		}

		// 检测到输入"bye"时，终止读取
		if string(line) == "bye" {
			return
		}
		ch <- line
	}
}

func write(conn net.Conn, ch <-chan []byte) {
	// 通道关闭时，关闭连接
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn close error:%v", err)
		}
	}()

	writer := bufio.NewWriter(conn)
	for msg := range ch {
		// 消息处理
		if _, err := writer.WriteString("you say: " + string(msg) + "\n"); err != nil {
			log.Printf("write error:%v", err)
		}
		if err := writer.Flush(); err != nil {
			log.Printf("flush error:%v", err)
		}
	}

	log.Println("bye")
}
