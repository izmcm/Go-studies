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
	commonServer()
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

}
