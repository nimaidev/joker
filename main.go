package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	isGracefulShutdown := make(chan bool)
	fmt.Println("Hello Joker")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a connection socket
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer listener.Close()

	go func() {
		for sig := range stop {
			<-stop
			fmt.Println("User Killing the process ", sig)
			isGracefulShutdown <- true
			listener.Close()
		}
	}()
	for {
		select {
		case <-isGracefulShutdown:
			fmt.Println("Killing listener")
			listener.Close()
			return
		default:
			con, err := listener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
					fmt.Println("Connection closed")
				default:
					fmt.Printf("Error %v \n", err)
				}
				return
			}
			go handleConnection(con, ctx) //To handle connection close
		}

	}
}

func handleConnection(con net.Conn, ctx context.Context) {
	fmt.Printf("Connection details %s \n", con.RemoteAddr())
	defer con.Close()
	// need to read what client are sending
	buffer := make([]byte, 1028)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Connection closed by server")
			return
		default:
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
}
