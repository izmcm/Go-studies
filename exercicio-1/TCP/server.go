package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running...")

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		// create thread
		go doConnection(conn)
	}
}

func doConnection(conn net.Conn) {
	fmt.Println("Connect with ", conn.RemoteAddr().String())

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Message from ", conn.RemoteAddr().String(), ": ", string(message))

		conn.Write([]byte(message + "\n"))
	}
}
