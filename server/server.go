package server

import (
	"encoding/binary"
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
	returnStr := "OK, " + string(data[4:]) //first four byte are tag and value

	tlv, err := Decode(data)
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
	log.Println("Size: ", len([]byte(tlv.Value)))
	// conn.Write([]byte(returnStr))
	log.Println("Return String:", returnStr)
	_, err = conn.Write([]byte(returnStr))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
}

type TLV struct {
	Tag    uint16 //2bit
	Length uint16 //bit
	Value  []byte //as per length
}

func Decode(data []byte) (*TLV, error) {
	//four bit are reserved for Key, and length
	if len(data) <= 5 {
		return nil, fmt.Errorf("insufficient data")
	}
	cmd := binary.BigEndian.Uint16(data[:2])
	length := binary.BigEndian.Uint16(data[2:4])

	if len(data) < int(length)+4 {
		return nil, fmt.Errorf("insufficient data for TLV value decoding")
	}

	nData := data[4:length]

	fmt.Println("DATA: ", string(nData), " CMD: ", cmd)

	tlv := TLV{
		Tag:    cmd,
		Length: length,
		Value:  nData,
	}

	return &tlv, nil
}

func (t *TLV) Encode() []byte {
	buf := make([]byte, 4+len(t.Value))
	binary.BigEndian.PutUint16(buf, t.Tag)
	binary.BigEndian.PutUint16(buf[2:], t.Length)
	copy(buf[4:], t.Value)
	// fmt.Println(buf)
	return buf
}
