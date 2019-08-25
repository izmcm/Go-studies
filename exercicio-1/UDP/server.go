package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server UDP is running...")

	addr, err := net.ResolveUDPAddr("udp4", "localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		doConnection(ln)
	}

}

func doConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := string(buffer[:n])
	fmt.Print("Received from ", addr, ": ", message)

	_, err = conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}
