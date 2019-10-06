package crh

import (
	"bufio"
	"fmt"
	"net"
)

type CRH struct {
	ServerHost string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

// TODO: ver se é isso mesmo
func (crh CRH) SendReceive(msgToServer []byte) []byte {
	addr := crh.ServerHost + ":" + string(crh.ServerPort)

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

	// send
	fmt.Fprintf(conn, string(msgToServer))

	// receive
	feedback, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	return []byte(feedback)
}
