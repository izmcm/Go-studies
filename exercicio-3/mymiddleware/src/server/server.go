package main

import (
	"crh"
	"fmt"
	"log"
	"miop"
	"net"
	"naming"
)

func main() {
	namingProxy := naming.NamingService{}

	calculator := proxies.NewCalculatorProxy()

	namingProxy.Register("Calculator", calculator)

	fmt.Println("Calculator server running!")
	calculatorInvoker := invoker.NewCalculatorInvoker()
	calculatorInvoker.Invoke()

	fmt.Scanln()
}
