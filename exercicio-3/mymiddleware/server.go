package main

import (
	// "crh"
	"fmt"
	// "log"
	// "miop"
	// "net"
	"clientproxy"
	"invoker"
	"naming"
	"proxies"
)

func main() {
	namingProxy := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}

	calculator := proxies.NewCalculatorProxy()
	namingProxy.Register("Calculator", calculator)

	// fmt.Println("Calculator server running!")
	calculatorInvoker := invoker.NewCalculatorInvoker()
	calculatorInvoker.Invoke()

	fmt.Scanln()
}
