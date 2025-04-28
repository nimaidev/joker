package main

import (
	"fmt"
	"net"
	"strings"
)

func main2() {
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
		go handleConnection(con) //To handle connection close

	}
}

func handleConnection(con net.Conn) {
	fmt.Printf("Connection details %s \n", con.RemoteAddr())
	defer con.Close()
	// need to read what client are sending
	buffer := make([]byte, 1028)
	for {
		n, err := con.Read(buffer)
		if err != nil {
			fmt.Printf("Error %v \n", err)
			return
		}
		clientBuffer := buffer[:n]
		clientCmd := string(clientBuffer)
		// Trim space from clientCmd
		clientCmd = strings.TrimSpace(clientCmd) //this will remove \n or \r
		if clientCmd == "0" {
			con.Write([]byte("Thank you for using JOKER"))
			con.Close()
			return
		}
		fmt.Printf("Client : %s | Data : %s", con.RemoteAddr(), string(clientBuffer))
		con.Write([]byte("OK\r\n"))

	}
}
