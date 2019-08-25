package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	lst := make([]time.Duration, 0, 0)

	conn, err := net.Dial("tcp", "127.0.0.1:8081") // connect to localhost
	if err != nil {
		fmt.Println(err)
		return
	}

	text := "seja muito bem vindo a vida real\n"

	for i := 0; i < 10000; i += 1 {
		// fmt.Print("Text to send: ")

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
		diff := tm2.Sub(tm1)
		lst = append(lst, diff)
		fmt.Println("delay time: ", diff)
		fmt.Print("Message from server: " + feedback)
	}

	// put in a csv
	f, err := os.Create("benchmark1.csv")
	if err != nil {
        panic(err)
    }
	defer f.Close()

	// fmt.Println(lst)
	for idx, tm := range lst {
		dt := fmt.Sprintf("%d,%s\n", idx, tm.String())
		// fmt.Println(dt)
		f.Write([]byte(dt))
	}
}
