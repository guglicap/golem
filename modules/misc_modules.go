package modules

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func padder(m Module) { //Hack-ish
	output <- Update{m.Position, m.Index, ""}
}

func syu(m Module) {
	lastCount := -45
	var count int
	if ok := inPATH("pacman"); ok != "" {
		output <- Update{m.Position, m.Index, ok}
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
			output <- Update{m.Position, m.Index, fmt.Sprintf("-Syu: %d", count)}
			lastCount = count
		}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
