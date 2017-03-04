package modules

import (
	"log"
	"net"
	"time"
)

func netAddress(m Module) {
	netInterface, err := net.InterfaceByName(m.options.NetInterface)
	if err != nil {
		errOutput(m, err)
		return
	}
	for {
		addrs, err := netInterface.Addrs()
		if err != nil {
			errOutput(m, err)
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
