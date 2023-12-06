package main

import (
	"fmt"
	"net"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello Joker")

	//create a tcp server
	lstn, err := net.Listen("tcp", "localhost:8080")

	handleError(err)

	fmt.Println("Server started:", lstn.Addr())
	for {
		con, err := lstn.Accept()
		handleError(err)
		handleConnection(con)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println(conn)
}
