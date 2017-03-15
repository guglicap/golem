package modules

import (
	"bufio"
	"errors"
	"os/exec"
	"regexp"
	"unicode"

	"github.com/guglicap/i3ipc"
)

func (m *WsModule) bspwm() {
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

func (m *WsModule) i3() {
	i3ipc.Init() //I know this sucks
	sock, err := i3ipc.GetIPCSocket()
	if err != nil {
		m.errOutput(err)
		return
	}
	events, err := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	if err != nil {
		m.errOutput(err)
		return
	}
	for {
		workspaces, err := sock.GetWorkspaces()
		if err != nil {
			m.errOutput(err)
			return
		}
		result := " "
		for _, w := range workspaces {
			if w.Focused {
				result += m.WsFocused
			} else {
				result += m.WsUnfocused
			}
			result += " "
		}
		output <- m.update(result)
		<-events //Block until a ws is focused
	}
}
