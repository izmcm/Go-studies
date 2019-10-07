package main

import (
	"fmt"
	// "invoker"
	// "log"
	"clientproxy"
	"naming"
	"proxies"
)

func main() {
	namingServer := naming.NamingService{Repository: make(map[string]clientproxy.ClientProxy)}
	namingServer.Register("Calculator", proxies.NewCalculatorProxy())
	// fmt.Println(namingServer.List())
	// calculator := namingServer.Lookup("Calculator").(proxies.CalculatorProxy)
	// calculator := proxies.NewCalculatorProxy()
	calculator := proxies.CalculatorProxy(namingServer.Lookup("Calculator"))

	fmt.Println(calculator.Add(1, 2))
}
