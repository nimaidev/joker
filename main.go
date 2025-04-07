package main

import (
	"fmt"

	"github.com/0x4E43/joker/server"
)

func main() {
	fmt.Println("Hello Joker")

	servOption := server.SetServerOption("6379")
	server.CreateServer(servOption)
}
