package shared

import (
// "fmt"
)

type Request struct {
	Op     string
	Params []interface{}
}

type Invocation struct {
	Host string
	Port int
	Request
}

const MIOP_REQUEST int = 1
const CALCULATOR_PORT int = 8081
const CALCULATOR_IP string = "localhost"
