package modules

import (
	"bufio"
	"os/exec"
	"regexp"
	"unicode"
)

func ws(m Module) {
	re := regexp.MustCompile("([oOuUfF]\\d)")
	if ok := inPATH("bspc"); ok != "" {
		output <- Update{m.slot, m.colors, ok}
		return
	}
	cmd := exec.Command("bspc", "subscribe")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		errOutput(m, err)
		return
	}
	scan := bufio.NewScanner(stdout)
	err = cmd.Start()
	result := " "
	for scan.Scan() {
		bspc := scan.Text()
		matches := re.FindAllStringSubmatch(bspc, -1)
		for _, match := range matches {
			if unicode.IsUpper(rune(match[1][0])) {
				result += m.options.WsFocused
			} else {
				result += m.options.WsUnfocused
			}
			result += " "
		}
		output <- Update{m.slot, m.colors, result}
		result = " "
	}

}
