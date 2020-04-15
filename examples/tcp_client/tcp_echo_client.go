package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	servAddr := "localhost:9010"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	//send message
	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	//receive buffer
	buf := make([]byte, 1024)

	//receive message
	_, err = conn.Read(buf[0:])

	if err != nil {
		fmt.Printf("Got error")
		os.Exit(1)
	}
	fmt.Printf("data received: %s", string(buf))

	//closing the connection
	conn.Close()
}
