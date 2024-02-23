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
	log.Println(str)
	if str == "PING\n" {
		return []byte("PONG")
	} else {
		return []byte("DUMB\n")
	}
	parts := strings.Split(str, constants.EOL)
	newPart := strings.Split(str, "\n")
	log.Println("PARTS: ", newPart[:1])
	if len(parts) > 0 {
		//read type by seeing the first character
		singleCmd := parts[0]
		if singleCmd != "" {
			cmdType := singleCmd[0]
			if string(cmdType) == constants.ARRAY {
				log.Println("Size: ", singleCmd[1:])
				return processCommand(parts[1 : len(parts)-1])
			}
		}
	}
	return respondError("Not enough parameter")
}

func processCommand(cmdList []string) []byte {
	log.Println("Inside Process command, size :", len(cmdList), cmdList)
	for i := 0; i < len(cmdList)-1; i++ {
		var cmdLen int
		if i%2 == 0 {
			cmdLen, _ = strconv.Atoi(cmdList[i][1:])
			if cmdLen != 0 {
				//check for the next value
				if cmdList[i+1] == constants.CLI_CMD_PING {
					return respondPingCommand("PING")
				}
			}
		}
		log.Println("Command length: ", cmdLen)
	}
	return respondPingCommand("dd")
}

func respondPingCommand(msg string) []byte {
	returnStr := constants.SIMPLE_STRING + "PONG" + constants.EOL
	//+ constants.SIZE + strconv.FormatInt(int64(len(msg)), 10) + constants.EOL + msg + constants.EOL
	log.Println(returnStr)
	return []byte(returnStr)
}

func respondError(msg string) []byte {
	return []byte(constants.ERROR + msg + constants.EOL)
}
