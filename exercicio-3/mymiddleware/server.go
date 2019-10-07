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
	lookupServer()
}

func commonServer() {
	namingProxy := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}

	calculator := proxies.NewCalculatorProxy()
	namingProxy.Register("Calculator", calculator)

	fmt.Println("Server up!")

	// fmt.Println("Calculator server running!")
	calculatorInvoker := invoker.NewCalculatorInvoker()
	// for {
	calculatorInvoker.Invoke()
	// }

	fmt.Scanln()
}

func lookupServer() {
	namingProxy := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	lookupInvoker := invoker.NewNameServerInvoker(namingProxy)

	go func() {
		lookupInvoker.Invoke()
	}()

	calculator := proxies.NewCalculatorProxy()
	namingProxy.Register("Calculator", calculator)

	fmt.Println("Server up!")

	calculatorInvoker := invoker.NewCalculatorInvoker()
	calculatorInvoker.Invoke()

	fmt.Scanln()
}
