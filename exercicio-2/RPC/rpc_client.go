package main

import (
	"encoding/csv"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"
)

type AddTask struct {
	Number1   int
	Number2   int
	Operation int
	Opid      int
}
type Calculator struct {
	Name string
}

var times [10000][]string

//cacula o tempo
func resolutionsWithTime(client *rpc.Client) int64 {
	var response int
	//v1 := [2]int{1, 2}
	v1 := AddTask{
		Number1:   1,
		Number2:   2,
		Operation: 0,
		Opid:      0,
	}

	//v2 := [2]int{4, 1}
	v2 := AddTask{
		Number1:   4,
		Number2:   1,
		Operation: 0,
		Opid:      0,
	}
	//v3 := [2]int{2, 5}
	v3 := AddTask{
		Number1:   2,
		Number2:   5,
		Operation: 0,
		Opid:      0,
	}
	//v4 := [2]int{4, 2}
	v4 := AddTask{
		Number1:   4,
		Number2:   2,
		Operation: 0,
		Opid:      0,
	}
	start := time.Now()
	client.Call("Calculator.Sum", v1, &response)
	log.Println(v1, response)
	client.Call("Calculator.Sub", v2, &response)
	log.Println(v2, response)
	client.Call("Calculator.Multiply", v3, &response)
	log.Println(v3, response)
	client.Call("Calculator.Divide", v4, &response)
	log.Println(v4, response)
	now := time.Now()
	timeAll, _ := time.ParseDuration(now.Sub(start).String())
	return timeAll.Nanoseconds()
}

func main() {
	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:8085")
	log.Println(client)
	var quant = 10000
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	for i := 0; i < quant; i++ {
		//times[i] = string{{strconv.FormatInt(resolutionsWithTime(client), 10)}, {"nano"}}
		times[i] = []string{strconv.FormatInt(resolutionsWithTime(client), 10)}
	}
	for i := 0; i < quant; i++ {
		log.Println(times[i])
	}
	creatFile()
}

func creatFileAndWrite() {

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range times {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
