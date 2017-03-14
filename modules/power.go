package modules

import (
	"encoding/json"
	"strings"
)

//PowerModule displays a set of button to poweroff, suspend or reboot the computer
type PowerModule struct {
	ModuleBase
	Format, PowerOffText, RebootText, SuspendText string
	PowerOffCmd, RebootCmd, SuspendCmd            string
}

//BuildPower initializes a PowerModule
func BuildPower(ms *ModuleSpec) Module {
	opts := struct {
		Format, PowerOffText, RebootText, SuspendText string
		PowerOffCmd, RebootCmd, SuspendCmd            string
	}{
		"%P",
		"\uf011",
		"\uf021",
		"\uf186",
		"poweroff",
		"reboot",
		"systemctl suspend",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &PowerModule{
		ModuleBase:   buildModuleBase(ms),
		Format:       opts.Format,
		PowerOffText: opts.PowerOffText,
		RebootText:   opts.RebootText,
		SuspendText:  opts.SuspendText,
		PowerOffCmd:  opts.PowerOffCmd,
		RebootCmd:    opts.RebootCmd,
		SuspendCmd:   opts.SuspendCmd,
	}
}

//Run starts the module
func (m *PowerModule) Run() {
	result := m.Format
	result = strings.Replace(result, "%P", buttonify(m.PowerOffCmd, m.PowerOffText), -1)
	result = strings.Replace(result, "%R", buttonify(m.RebootCmd, m.RebootText), -1)
	result = strings.Replace(result, "%S", buttonify(m.SuspendCmd, m.SuspendText), -1)
	output <- m.update(result)
}
