package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

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
		// scanner := bufio.NewScanner(conn)
		// rcv := scanner.Text()
		// conn.Write([]byte(rcv))
		// for scanner.Scan() {
		// 	rcv := scanner.Text()
		// 	conn.Write([]byte(rcv))
		// }hh
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
		log.Println("Read", n, "bytes from connection: ", conn.RemoteAddr())
		// log.Println(string(buf))
		// log.Println("BUFFER: ", buf)
		// conn.Write(parser.ProcessCLICommand(string(buf)))
		return
		conn.Write(buf)
		buf = nil
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var commandBuffer strings.Builder
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		recv := scanner.Text()
		// returnStr := recv + "\r\n"
		commandBuffer.WriteString(recv + " ")

		log.Println("CLIENT:", conn.RemoteAddr(), ":", recv)
		// fmt.Println("SERVER:", returnStr)
		if len(recv) > 1 {
			parseCommand(recv)
		}

		if recv == "" {
			// Process the entire command
			command := strings.TrimSpace(commandBuffer.String())
			if len(command) > 0 {
				parseCommand(command)
			}

			// Clear the buffer for the next command
			commandBuffer.Reset()
		}
		conn.Write([]byte(recv))
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

func parseCommand(cmd string) {
	parts := strings.Fields(cmd)
	cmdLen := len(parts)
	if cmdLen < 1 {
		fmt.Println("Invalid Command")
		return
	}

	fmt.Println(cmdLen, "arguments:")
	for i, arg := range parts {
		fmt.Printf("%d) \"%s\"\n", i+1, arg)
	}
	wholeCommand := strings.Join(parts, " ")
	fmt.Println("Entire Command as a Single Argument:")
	fmt.Printf("1) \"%s\"\n", wholeCommand)
}
