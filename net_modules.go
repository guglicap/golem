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
	defaultInterface *net.Interface
)

func getDefaultInterface() {
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) < 1 {
		defaultInterface = nil
		return
	}
	for _, i := range interfaces {
		if !strings.Contains(i.Flags.String(), "LOOPBACK") { //Making assumptions here.
			defaultInterface = &i
		}
	}
}

func netAddr(m Module) {
	runOnce := checkDuration(m.Refresh)
	if defaultInterface == nil {
		return
	}
	for {
		addrs, err := defaultInterface.Addrs()
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

func ping(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		cmd, err := exec.Command("ping", "-c 1", "google.com").Output()
		if err != nil {
			log.Println(err)
			out <- Update{m.position, m.index, errorColor("Can't ping shit " + err.Error())}
		}
		re := regexp.MustCompile("time=(\\d+\\.\\d+ ms)")
		matches := re.FindStringSubmatch(string(cmd))
		if len(matches) < 2 {
			log.Println(matches)
			out <- Update{m.position, m.index, errorColor("Can't ping ")}
			return
		}
		out <- Update{m.position, m.index, "google.com: " + matches[1]}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}
