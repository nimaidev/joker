package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello Joker")

	// create a connection socket
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer listener.Close()
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error %v \n", err)
			return
		}
		defer con.Close()
		fmt.Printf("Connection details %s", con.RemoteAddr())

	}
}
