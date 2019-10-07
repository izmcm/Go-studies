package srh

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type SRH struct {
	ServerHost string
	ServerPort int
	Conn       *net.Conn
	Listenner  *net.Listener
}

var ln net.Listener
var conn net.Conn
var err error

// TODO: ver se é isso mesmo

func (srh *SRH) Receive() []byte {
	addr := srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
	fmt.Printf("listening to addres: ")
	fmt.Println(addr)

	if srh.Conn == nil {
		ln, err = net.Listen("tcp", addr)
		if err != nil {
			fmt.Println(err)
			return []byte("error")
		}
		srh.Listenner = &ln
	} else {
		ln = *srh.Listenner
	}

	// fmt.Println("Calculator server running in", addr)
	conn, err = ln.Accept()
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	fmt.Println("connect with", conn)
	srh.Conn = &conn
	// fmt.Printf("tipo: %T\n", conn)

	// receive
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []byte("error")
	}

	// fmt.Print("Message from ", conn.RemoteAddr().String(), ": ", string(message))
	return []byte(message)
}

func (srh *SRH) Send(msgToClient []byte) {
	// fmt.Println("\n\n\nConnection")
	// fmt.Println(conn)
	// fmt.Println(srh.Conn)
	fmt.Println("vai enviar")
	fmt.Fprintf(conn, string(msgToClient)+"\n")
	fmt.Fprintf(*srh.Conn, string(msgToClient)+"\n")
	fmt.Println("enviou")

	(*srh.Conn).Close()
}
