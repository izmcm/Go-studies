package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Print("Text to send: ")
		reader := bufio.NewReader(os.Stdin)

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := []byte(text)

		// send
		_, err = conn.Write(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// receive
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message from server: ", string(buffer[:n]))
	}
}
