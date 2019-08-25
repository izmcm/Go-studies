package main

import (
	// "bufio"
	"fmt"
	"net"
	"os"
	"time"
	"sync"
)

var wg sync.WaitGroup

func benchmark(text string, iterations int, num int, total int) {
	lst := make([]time.Duration, 0, 0)

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

	defer conn.Close()

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
		_, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		tm2 := time.Now()
		diff := tm2.Sub(tm1)
		lst = append(lst, diff)
		fmt.Println("delay time for ", num, ": ", diff)
		// fmt.Println("Message from server: ", string(buffer[:n]))
	}

	// put in a csv
	name := fmt.Sprintf("UDP/benchmark_%d/benchmark_tcp_%d.csv", total, num)
	f, err := os.Create(name)
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


	defer wg.Done()
}

func main() {

	text := "seja muito bem vindo a vida real\n"
	iterations := 10000
	num := 6

	wg.Add(num)

	for i := 0; i < num; i += 1 {
		go benchmark(text, iterations, i, num)
	}

	wg.Wait()
}
