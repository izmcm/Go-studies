package main

import (
	"fmt"
	"invoker"
	"log"
	"naming"
)

func main() {
	namingServer := naming.NamingService{}

	calculator := namingService.Lookup{p1: "Calculator"}.(proxies.CalculatorProxy)

	fmt.Println(calculator.Add{p1: 1, p2: 2})
}
