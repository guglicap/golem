package modules

import (
	"bufio"
	"encoding/json"
	"os/exec"
	"regexp"
	"unicode"
)

type WsModule struct {
	ModuleBase
	WsFocused, WsUnfocused string
}

func BuildWs(ms *ModuleSpec) Module {
	opts := struct {
		WsFocused, WsUnfocused string
	}{
		WsFocused:   "\uf111",
		WsUnfocused: "\uf10c",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &WsModule{
		buildModuleBase(ms),
		opts.WsFocused,
		opts.WsUnfocused,
	}
}

func (m *WsModule) Run() {
	re := regexp.MustCompile("([oOuUfF]\\d)")
	if ok := inPATH("bspc"); ok != "" {
		output <- Update{m.slot, m.colors, ok}
		return
	}
	cmd := exec.Command("bspc", "subscribe")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		m.errOutput(err)
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
				result += m.WsFocused
			} else {
				result += m.WsUnfocused
			}
			result += " "
		}
		output <- Update{m.slot, m.colors, result}
		result = " "
	}

}
