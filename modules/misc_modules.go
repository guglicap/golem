package modules

import (
	"errors"
	"strings"
)

func text(m Module) {
	output <- Update{m.slot, m.colors, m.options.Text}
}

func button(m Module) {
	cmd, txt := m.options.Command, m.options.ButtonText
	if len(cmd) == 0 {
		errOutput(m, errors.New("You didn't specify a command."))
		return
	}
	if len(txt) == 0 {
		txt = cmd
	}
	if ok := inPATH(cmd); ok != "" {
		output <- Update{m.slot, m.colors, ok}
		return
	}
	output <- Update{m.slot, m.colors, buttonify(cmd, txt)}
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
	output <- Update{m.slot, m.colors, result}
}
