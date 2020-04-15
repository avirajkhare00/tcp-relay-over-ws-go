package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:9010"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		fmt.Printf("Unable to resolve tcpAddr: %s", tcpAddr)
	}

	//listen for incoming connection
	listener, err := net.ListenTCP("tcp", tcpAddr)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Unable to accept: %s", err)
		os.Exit(1)
	}

	incomingMsg := make([]byte, 1024)

	//receive message
	_, err = conn.Read(incomingMsg)
	if err != nil {
		fmt.Printf("Unable to read message %s", err)
		os.Exit(1)
	}

	//send message back to client
	_, err = conn.Write(incomingMsg)
	if err != nil {
		fmt.Printf("Unable to send message to client: %s", err)
		os.Exit(1)
	}
}
