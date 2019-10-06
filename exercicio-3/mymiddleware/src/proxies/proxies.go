package proxies

import (
	"clientProxy"
	"requestor"
	"shared"
)

type CalculatorProxy clientProxy.ClientProxy

func NewCalculatorProxy() CalculatorProxy {
	// TODO
}

func (proxy CalculatorProxy) Add(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2
	// Requestor preparation
	request := shared.Request{Op: "Add", Params: params}
	inv := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Sub(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Sub", Params: params}
	inv := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Mul(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Mul", Params: params}
	inv := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}

func (proxy CalculatorProxy) Div(p1 int, p2 int) int {
	// Client info
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Requestor preparation
	request := shared.Request{Op: "Div", Params: params}
	inv := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}
