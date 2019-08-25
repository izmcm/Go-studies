package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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
		tm1 := time.Now()
		fmt.Fprintf(conn, text+"\n") // send to server

		// receive
		feedback, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		tm2 := time.Now()
		fmt.Println("delay time: ", tm2.Sub(tm1))
		fmt.Print("Message from server: " + feedback)
	}
}
