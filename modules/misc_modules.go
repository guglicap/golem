package modules

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func text(m Module) {
	output <- Update{m.position, m.index, m.options.Text}
}

func button(m Module) {
	cmd, txt := m.options.Command, m.options.ButtonText
	if len(cmd) == 0 {
		output <- Update{m.position, m.index, colorize(errorColor, "You didn't specify a command.")}
	}
	if len(txt) == 0 {
		txt = cmd
	}
	if ok := inPATH(cmd); ok != "" {
		output <- Update{m.position, m.index, ok}
		return
	}
	output <- Update{m.position, m.index, buttonify(cmd, txt)}
}

func icontray(m Module) {
	cmds := strings.Split(m.options.IconTrayCommands, ",")
	txts := strings.Split(m.options.IconTrayText, ",")
	result := " "
	for i, cmd := range cmds {
		if len(txts) <= i {
			result += buttonify(cmd, cmd+" ")
		} else {
			result += buttonify(cmd, txts[i]+" ")
		}
	}
	output <- Update{m.position, m.index, result}
}
