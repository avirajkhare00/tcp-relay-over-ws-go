package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	authToken := strings.Trim(authHeader, "Bearer ")

	someToken := "ABCD1234"

	if authToken != someToken {
		w.WriteHeader(401)
		w.Write([]byte("Not authorized"))
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("upgrade: %s\n", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("read: %s\n", err)
			break
		}
		fmt.Printf("recv: %s\n", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			fmt.Printf("write: %s\n", err)
			break
		}
	}
}

func main() {
	fmt.Println("Let's go!")

	serverAddress := "localhost:8081"

	http.HandleFunc("/echo", echo)
	http.ListenAndServe(serverAddress, nil)
}
