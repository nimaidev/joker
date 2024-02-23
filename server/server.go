package server

import (
	"fmt"
	"io"
	"log"
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
	defer lstnr.Close()
	for {
		conn, err := lstnr.Accept()
		log.Println("Client: ", conn.RemoteAddr())
		utils.HandleError(err)
		go handleConnectionV0(conn)
	}
}

func handleConnectionV0(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client, ", conn.RemoteAddr())
			}

			break
		}
		if n == 0 {
			//no data to process
			continue
		}
		log.Println("Read", n, "bytes from connection: ", conn.RemoteAddr())
		// Process the received data
		processData(buf[:n], conn)

	}
}

func processData(data []byte, conn net.Conn) {
	// Process the received data here
	returnStr := "OK, " + string(data[4:]) //first four bute are tag and value
	log.Println("Size: ", len([]byte(returnStr)))
	// conn.Write([]byte(returnStr))
	log.Println("Return String:", returnStr)
	_, err := conn.Write([]byte(returnStr))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
}
