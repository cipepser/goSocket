package main

import (
	"fmt"
	"net"
	"time"
	"log"
)

var (
	port = "8080" // listen port
)

func handleClient(conn net.Conn)  {
	defer conn.Close()
	fmt.Println("client accept!")
	
	// recieve messege
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	messageBuf := make([]byte, 1024)
	massegeLen, err := conn.Read(messageBuf)
	if err != nil {
		log.Fatal(err)
	}	
	message := string(messageBuf[:massegeLen])
	
	// send messeage
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))
}

func main() {
	// port監視
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":" + port) // end point of L4
	if err != nil {
		log.Fatal(err)
	}
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	
	// 新規接続があればgoroutine起動
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go handleClient(conn)
	}
}