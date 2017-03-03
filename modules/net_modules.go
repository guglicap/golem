package modules

import (
	"log"
	"net"
	"strings"
	"time"
)

var (
	defaultNetInterface *net.Interface
)

func getDefaultInterface() {
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) < 1 {
		defaultNetInterface = nil
		return
	}
	for _, i := range interfaces {
		if !strings.Contains(i.Flags.String(), "LOOPBACK") { //Making assumptions here.
			defaultNetInterface = &i
		}
	}
}

func netAddress(m Module) {
	netInterface, err := net.InterfaceByName(m.options.NetInterface)
	if err != nil {
		if defaultNetInterface != nil {
			netInterface = defaultNetInterface
		} else {
			return
		}
	}
	for {
		addrs, err := netInterface.Addrs()
		if err != nil {
			errOutput(m, err)
			return
		}
		output <- Update{m.position, m.index, addrs[0].String()}
		log.Println(addrs)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
