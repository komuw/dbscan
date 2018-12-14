package main

import (
	"log"
	"net"
)

/*
usage:
  go run -race reverse.go
  echo -n "test out the server" | nc localhost 3333
  curl -vkIL localhost:3333
*/

const (
	connHost = "localhost"
	connPort = "3333"
	connType = "tcp"
)

func main() {
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	defer l.Close()

	log.Println("Listening on " + connHost + ":" + connPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting:", err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// TODO: make the buffer growable
	buf := make([]byte, 100)
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading:", err)
	}
	_ = reqLen

	log.Println("received::", buf)
	log.Println("received2::", string(buf))

	conn.Write([]byte("Message received."))
}