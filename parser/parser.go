package parser

import (
	"log"
	"strconv"
	"strings"

	"github.com/0x4E43/joker/constants"
)

// TODO: Step 1 Parse from redis cli
func ProcessCLICommand(str string) []byte {
	log.Println("Function called with ", str)
	str = strings.Trim(str, " ")
	parts := strings.Split(str, constants.EOL)
	log.Println(parts)
	if len(parts) > 0 {
		//read type by seeing the first character
		singleCmd := parts[0]
		cmdType := singleCmd[0]
		if string(cmdType) == constants.ARRAY {
			log.Println("Size: ", singleCmd[1:])
			return processCommand(parts[1:])
		}
	}
	return respondError("nOT ENOUGH PARAMETER")
}

func processCommand(cmdList []string) []byte {
	log.Println("Inside Process command, size :", len(cmdList), cmdList)
	for i := 0; i < len(cmdList)-1; i++ {
		var cmdLen int
		if i%2 == 0 {
			cmdLen, _ = strconv.Atoi(cmdList[i][1:])
		}
		log.Println("Command length: ", cmdLen)
	}
	return respondPingCommand()
}

func respondPingCommand() []byte {
	return []byte(constants.SIMPLE_STRING + "PONG" + constants.EOL)
}

func respondError(msg string) []byte {
	return []byte(constants.ERROR + msg + constants.EOL)
}
