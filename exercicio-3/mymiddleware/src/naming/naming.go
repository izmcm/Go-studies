package naming

import (
	"clientproxy"
	// "proxies"
	// "fmt"
)

type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

func (naming *NamingService) Register(name string, proxy clientproxy.ClientProxy) {
	naming.Repository[name] = proxy
}

func (naming NamingService) Lookup(name string) clientproxy.ClientProxy {
	return naming.Repository[name]
}

func (naming NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Repository
}

// TODO: Unbind / Rebind
