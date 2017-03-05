package modules

import (
	"encoding/json"
	"strings"
)

type LauncherModule struct {
	ModuleBase
	Programs, Icons string
}

func BuildLauncher(ms *ModuleSpec) Module {
	opts := struct {
		Programs, Icons string
	}{
		"firefox",
		"\uf269",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &LauncherModule{
		buildModuleBase(ms),
		opts.Programs,
		opts.Icons,
	}
}

func (m *LauncherModule) Run() {
	cmds := strings.Split(m.Programs, ",")
	txts := strings.Split(m.Icons, ",")
	result := " "
	for i, cmd := range cmds {
		if len(txts) <= i || txts[i] == "_" {
			result += buttonify(cmd, cmd+" ")
		} else {
			result += buttonify(cmd, txts[i]+" ")
		}
	}
	output <- Update{m.slot, m.colors, result}
}
