package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081") // connect to localhost
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

		// send
		fmt.Fprintf(conn, text+"\n") // send to server

		// receive
		feedback, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("Message from server: " + feedback)
	}
}
