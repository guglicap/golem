package main

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

func NetAddr(m Module) {
	runOnce := checkDuration(m.Refresh)
	if defaultNetInterface == nil {
		return
	}
	for {
		addrs, err := defaultNetInterface.Addrs()
		if err != nil {
			out <- Update{m.position, m.index, errorColor(err.Error())}
		}
		out <- Update{m.position, m.index, addrs[0].String()}
		log.Println(addrs)
		if runOnce {
			log.Println("runOnce true, exiting.")
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}

func Ping(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		cmd, err := exec.Command("ping", "-c 1", options.PingAddr).Output()
		if err != nil {
			log.Println(err)
			out <- Update{m.position, m.index, errorColor("Can't ping " + options.PingAddr + ", " + err.Error())}
		}
		re := regexp.MustCompile("time=(\\d+\\.\\d+ ms)")
		matches := re.FindStringSubmatch(string(cmd))
		if len(matches) < 2 {
			log.Println(matches)
			out <- Update{m.position, m.index, errorColor("Something went wrong pinging " + options.PingAddr)}
			return
		}
		out <- Update{m.position, m.index, options.PingAddr + ": " + matches[1]}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}
