package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
	"sync"
)

var wg sync.WaitGroup

// func genTimes(text string, iterations int, num int, total int, connection) {
func genTimes(text string, iterations int, num int, total int) {
	lst := make([]time.Duration, 0, 0)
	fmt.Println(num)

	connection, err := net.Dial("tcp", "127.0.0.1:8082") // connect to localhost
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	for i := 0; i < iterations; i += 1 {
		// fmt.Print("Text to send: ")

		// send
		tm1 := time.Now()
		fmt.Fprintf(connection, text+"\n") // send to server

		// receive
		_, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		tm2 := time.Now()
		diff := tm2.Sub(tm1)
		lst = append(lst, diff)
		fmt.Println("delay time for ", num, ": ", diff)
		// fmt.Print("Message from server: " + feedback)
	}

	// put in a csv
	name := fmt.Sprintf("TCP/benchmark_%d/benchmark_tcp_%d.csv", total, num)
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

	// lst := make([]time.Duration, 0, 0)

	text := "seja muito bem vindo a vida real\n"
	iterations := 10000
	num := 6

	wg.Add(num)

	for i := 0; i < num; i += 1 {
		go genTimes(text, iterations, i, num)
	}

	wg.Wait()
}
