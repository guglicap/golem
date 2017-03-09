package modules

import (
	"bufio"
	"encoding/json"
	"errors"
	"os/exec"
	"regexp"
	"unicode"
)

//WsModule is a workspace indicator for bspwm.
type WsModule struct {
	ModuleBase
	WsFocused, WsUnfocused string
}

//BuildWs initializes a WsModule
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

//Run starts the module.
func (m *WsModule) Run() {
	re := regexp.MustCompile("([oOuUfF]\\d)")
	if ok := inPATH("bspc"); !ok {
		m.errOutput(errors.New("Can't find bspc in $PATH"))
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
		output <- m.update(result)
		result = " "
	}

}
