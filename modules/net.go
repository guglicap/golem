package modules

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type NetModule struct {
	ModuleBase
	Interface string
}

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
		output <- Update{m.slot, m.colors, addrs[0].String()}
		log.Println(addrs)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
