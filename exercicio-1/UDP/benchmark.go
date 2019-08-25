package main

import (
	// "bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func benchmark() {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10000; i += 1 {
		// fmt.Print("Text to send: ")
		// reader := bufio.NewReader(os.Stdin)

		// text, err := reader.ReadString('\n')
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		message := []byte(text)

		// send
		tm1 := time.Now()
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

		tm2 := time.Now()
		diff := tm2.Sub(tm1)
		lst = append(lst, diff)
		fmt.Println("delay time: ", diff)
		fmt.Println("Message from server: ", string(buffer[:n]))
	}
}

func main() {
	lst := make([]time.Duration, 0, 0)

	addr, err := net.ResolveUDPAddr("udp", "localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	text := "seja muito bem vindo a vida real\n"


	// put in a csv
	f, err := os.Create("benchmark_udp_1.csv")
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
