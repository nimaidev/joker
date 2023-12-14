package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/0x4E43/joker/constants"
	"github.com/0x4E43/joker/utils"
)

type ServerOption struct {
	port string
}

func SetServerOption(port string) *ServerOption {
	opt := ServerOption{port}
	return &opt
}

func CreateServer(servOption *ServerOption) {
	//Listen to the port
	lstnr, err := net.Listen("tcp", ":"+servOption.port)
	utils.HandleError(err)
	fmt.Println("Joker laighing at :", lstnr.Addr().String())
	for {
		conn, err := lstnr.Accept()
		utils.HandleError(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		recv := scanner.Text()
		returnStr := constants.String + "SERVER: " + recv + constants.EOL
		fmt.Println("CLIENT: ", conn.RemoteAddr(), " : ", recv)
		conn.Write([]byte(returnStr))
	}

	if err := scanner.Err(); err != nil {
		fmt.Print("Something went wrong with Scanner")
	}
}
