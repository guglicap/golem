package modules

import (
	"log"
	"net"
	"os/exec"
	"regexp"
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

func netAddr(m Module) {
	if defaultNetInterface == nil {
		return
	}
	for {
		addrs, err := defaultNetInterface.Addrs()
		if err != nil {
			output <- Update{m.Position, m.Index, colorize(errorColor, err.Error())}
		}
		output <- Update{m.Position, m.Index, addrs[0].String()}
		log.Println(addrs)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}

func ping(m Module) {
	addr := m.options.PingAddress
	cmd, err := exec.Command("ping", "-c 1", addr).Output()
	if err != nil {
		log.Println(err)
		output <- Update{m.Position, m.Index, colorize(errorColor, "Are you sure "+addr+" is reachable?")}
		return
	}
	for {
		re := regexp.MustCompile("time=(\\d+\\.\\d+ ms)")
		matches := re.FindStringSubmatch(string(cmd))
		if len(matches) < 2 {
			log.Println(matches)
			output <- Update{m.Position, m.Index, colorize(errorColor, "Can't ping "+addr)}
		}
		output <- Update{m.Position, m.Index, addr + ": " + matches[1]}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
