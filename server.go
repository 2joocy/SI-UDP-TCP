package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	defer listener.Close()
	fmt.Println("Started server on 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	output := fmt.Sprintf("Received Message: " + message)
	fmt.Println(output)
	conn.Write([]byte(message))
	conn.Close()
}
