package modules

import (
	"encoding/json"
	"net"
	"time"
)

//NetModule displays the machine IP Address
type NetModule struct {
	ModuleBase
	Interface string
}

//BuildNet initializes a NetModule
func BuildNet(ms *ModuleSpec) Module {
	opts := struct {
		NetInterface string
	}{
		"enp3s0",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &NetModule{
		buildModuleBase(ms),
		opts.NetInterface,
	}
}

//Run starts the module
func (m *NetModule) Run() {
	netInterface, err := net.InterfaceByName(m.Interface)
	if err != nil {
		m.errOutput(err)
		return
	}
	for {
		addrs, err := netInterface.Addrs()
		if err != nil {
			m.errOutput(err)
			return
		}
		output <- m.update(addrs[0].String())
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
