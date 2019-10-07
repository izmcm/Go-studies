package crh

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
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
	addr := crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)

	conn, err := net.Dial("tcp", addr) // connect to localhost
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}
	fmt.Println("Calculator client running in", addr)
	fmt.Println(conn)

	// send
	fmt.Fprintf(conn, string(msgToServer)+"\n")
	fmt.Println("send", string(msgToServer), "to server")

	// receive
	feedback, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("-------- deu ruim --------")
		fmt.Println(err)
		return []byte("error")
	}
	fmt.Println("-------- deu bom --------")

	// lst := make([]byte, 3)
	// return lst
	return []byte(feedback)

}
