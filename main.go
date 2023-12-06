package main

import (
	"bufio"
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
	defer conn.Close()
	fmt.Println("Client connected: ", conn.RemoteAddr())

	//read data from client
	scanner := bufio.NewScanner(conn) //Initializing scanner at conn

	for scanner.Scan() {
		recieved := scanner.Text()
		fmt.Println("Recieved: ", recieved)
		//Write into the client
		conn.Write([]byte("Hello Client from: \n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in reading data: ", err)
	}

	fmt.Println("Client disconnected", conn.RemoteAddr())
}
