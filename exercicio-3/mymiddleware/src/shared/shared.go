package shared

import (
// "fmt"
)

type Invocation struct {
	Host string
	Port int
	Request
}

type Request struct {
	Op     string
	Params []interface{}
}

const MIOP_REQUEST int = 1
