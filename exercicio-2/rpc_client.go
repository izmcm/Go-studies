package main

import (
	"log"
	"net/rpc"
)

type Calculator struct {
	Name string
}

func main() {
	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	log.Println(client)
	var response int

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	v1 := [2]int{1, 2}
	v2 := [2]int{4, 1}
	v3 := [2]int{2, 5}
	v4 := [2]int{4, 2}

	client.Call("Calculator.Sum", v1, &response)
	log.Println(v1, response)
	client.Call("Calculator.Sub", v2, &response)
	log.Println(v2, response)
	client.Call("Calculator.Multiply", v3, &response)
	log.Println(v3, response)
	client.Call("Calculator.Divide", v4, &response)
	log.Println(v4, response)
}
