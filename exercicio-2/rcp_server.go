package main

// The net/rpc package stipulates that only methods that satisfy the
// following criteria will be made available for remote access; other methods will be ignored.

// The method’s type is exported. or builtin [OK]
// The method is exported [OK]
// The method has two arguments, both exported (or builtin types). [X]
// The method’s second argument is a pointer [X]
// The method has return type error [X]

// ˆˆˆˆˆˆˆˆˆˆˆˆ

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Calculator struct {
	Name string
}

func (c *Calculator) Sum(elements [2]int, result *int) error {
	*result = elements[0] + elements[1]
	return nil
}

func (c *Calculator) Sub(elements [2]int, result *int) error {
	*result = elements[0] - elements[1]
	return nil
}

func (c *Calculator) Multiply(elements [2]int, result *int) error {
	*result = elements[0] * elements[1]
	return nil
}

func (c *Calculator) Divide(elements [2]int, result *int) error {
	*result = elements[0] / elements[1]
	return nil
}

func testFunctions() {
	c := new(Calculator)
	var result int
	log.Println("1. Sum: 1 + 2 = ", c.Sum([2]int{1, 2}, &result))
	log.Println("> ", result)
	log.Println("2. Sub: 4 - 1 = ", c.Sub([2]int{4, 1}, &result))
	log.Println("> ", result)
	log.Println("3. Multiply: 2*5 = ", c.Multiply([2]int{2, 5}, &result))
	log.Println("> ", result)
	log.Println("1. Divide: 4/2 = ", c.Divide([2]int{4, 2}, &result))
	log.Println("> ", result)
}

// Bora dale kk
func main() {
	calc := new(Calculator)
	// var result int
	// Publish the receivers methods
	err := rpc.Register(calc)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 8081)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
