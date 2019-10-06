package srh

import (
	"bufio"
	"fmt"
	"net"
)

type SRH struct {
	ServerHost string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

// TODO: ver se é isso mesmo

func (srh SRH) Receive() []byte {
	addr := srh.ServerHost + ":" + string(srh.ServerPort)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	// receive
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	// fmt.Print("Message from ", conn.RemoteAddr().String(), ": ", string(message))
	return []byte(message)
}

func (SRH) Send(msgToClient []byte) {
	conn.Write([]byte(msgToClient))
}
