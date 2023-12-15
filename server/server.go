package server

import (
	"bufio"
	"fmt"
	"net"

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
		returnStr := recv + "\r\n"

		fmt.Println("CLIENT:", conn.RemoteAddr(), ":", recv)
		fmt.Println("SERVER:", returnStr)

		conn.Write([]byte(returnStr))
	}

	if err := scanner.Err(); err != nil {
		if opErr, ok := err.(*net.OpError); ok {
			if opErr.Err.Error() == "use of closed network connection" {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Network error:", opErr.Err)
			}
		} else {
			fmt.Println("Scanner error:", err)
		}
	}

}
