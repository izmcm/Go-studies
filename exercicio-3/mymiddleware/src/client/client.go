package main

import (
	"fmt"
	// "invoker"
	// "log"
	"naming"
	"proxies"
)

func main() {
	namingServer := naming.NamingService{}

	// calculator := namingServer.Lookup("Calculator").(proxies.CalculatorProxy)
	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))

	fmt.Println(calculator.Add(1, 2))
}
