package naming

import (
	"clientProxy"
	"proxies"
	// "fmt"
)

type NamingService struct {
	Repository map[string]proxies.CalculatorProxy
}

func (naming *NamingService) Register(name string, proxy clientProxy.ClientProxy) bool {
	naming.Repository[name] = proxy
}

func (naming NamingService) Lookup(name string) clientProxy.ClientProxy {
	return naming.Repository[name]
}

func (naming NamingService) List(name string) map[string]clientProxy.ClientProxy {}

// TODO: Unbind / Rebind
