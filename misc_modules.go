package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func padder(m Module) { //Hack-ish
	out <- Update{m.position, m.index, ""}
}

func syu(m Module) {
	lastCount := -45
	var count int
	runOnce := checkDuration(m.Refresh)
	if ok := inPATH("pacman"); ok != "" {
		out <- Update{m.position, m.index, ok}
		return
	}
	for {
		cmd, err := exec.Command("pacman", "-Qu").Output()
		if err != nil { //pacman will exit with status 1 when there are no updates.
			count = 0
			log.Println("Pacman err", err)
		} else {
			count = len(strings.Split(string(cmd), "\n")) - 1
		}
		log.Println("pacman output:", string(cmd))
		if count != lastCount {
			out <- Update{m.position, m.index, fmt.Sprintf("-Syu: %d", count)}
			lastCount = count
		}
		if runOnce {
			return
		}
		time.Sleep(time.Duration(m.Refresh))
	}
}
