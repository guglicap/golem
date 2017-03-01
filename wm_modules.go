package main

import (
	"os/exec"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func Ws(m Module) {
	lastActive := -12
	re := regexp.MustCompile("([oOfF]\\d)")
	runOnce := checkDuration(m.Refresh)
	for {
		cmd, err := exec.Command("bspc", "wm", "-g").Output()
		if err != nil {
			continue
		}
		matches := re.FindAllStringSubmatch(string(cmd), -1)
		if matches == nil {
			continue
		}
		var spaces = make([]string, 0)
		var active int
		for i, m := range matches {
			if unicode.IsUpper(rune(m[1][0])) {
				spaces = append(spaces, options.WsFocused)
				active = i
			} else {
				spaces = append(spaces, options.WsUnfocused)
			}
		}
		if active != lastActive {
			out <- Update{m.position, m.index, strings.Join(spaces, " ")}
			lastActive = active
		}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}
