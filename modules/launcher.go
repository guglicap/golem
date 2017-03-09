package modules

import (
	"encoding/json"
	"strings"
)

//LauncherModule displays an icon tray
type LauncherModule struct {
	ModuleBase
	Programs, Icons string
}

//BuildLauncher initializes a LauncherModule
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

//Run starts the module
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
	output <- m.update(result)
}
