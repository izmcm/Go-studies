package proxies

import (
	"clientproxy"
	"fmt"
	"requestor"
	"shared"
)

type CalculatorProxy clientproxy.ClientProxy

func NewCalculatorProxy() clientproxy.ClientProxy {
	// TODO: fazer uma inicialização com id arbitrário
	cp := clientproxy.ClientProxy{Host: shared.CALCULATOR_IP, Port: shared.CALCULATOR_PORT, Id: 1, TypeName: "Calculadora"}
	return cp
}

func (proxy CalculatorProxy) Add(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Add", Params: params}
	inv := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv)

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Sub(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Sub", Params: params}
	inv := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv)

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Mul(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Mul", Params: params}
	inv := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv)

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Div(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Div", Params: params}
	inv := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv)

	return int(ter[0].(float64))
}

// Zone for LookupProxy

type LookupProxy clientproxy.ClientProxy

func NewLookupProxy() LookupProxy {
	// TODO: fazer uma inicialização com id arbitrário
	cp := LookupProxy{Host: shared.NAMESERVER_IP, Port: shared.NAMESERVER_PORT, Id: 1, TypeName: "LookupTable"}
	return cp
}

func (proxy LookupProxy) Lookup(name string) clientproxy.ClientProxy {
	// Client info
	params := make([]interface{}, 1)
	params[0] = name

	// Requestor preparation
	request := shared.Request{Op: "GetServer", Params: params}
	inv := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv)
	fmt.Printf("recebido o dado: ")
	fmt.Println(ter)

	res := ter[0].(map[string]interface{})
	fmt.Printf("\ntype: %T\n", res)
	fmt.Println("value:\n", res)
	h := res["Host"].(string)
	p := int(res["Port"].(float64))
	i := int(res["Id"].(float64))
	t := res["TypeName"].(string)
	data := clientproxy.ClientProxy{Host: h, Port: p, Id: i, TypeName: t}

	fmt.Printf("\ntype: %T\n", data)
	fmt.Println("value:\n", data)

	return data
	// return ter[0].(map[string]interface{}).(map[string]clientproxy.ClientProxy)
}
