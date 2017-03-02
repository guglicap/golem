package modules

import (
	"os/exec"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func ws(m Module) {
	lastActive := -12
	re := regexp.MustCompile("([oOfF]\\d)")
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
		for i, match := range matches {
			if unicode.IsUpper(rune(match[1][0])) {
				spaces = append(spaces, m.options.WsFocused)
				active = i
			} else {
				spaces = append(spaces, m.options.WsUnfocused)
			}
		}
		if active != lastActive {
			output <- Update{m.Position, m.Index, strings.Join(spaces, " ")}
			lastActive = active
		}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
