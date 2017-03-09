package modules

import (
	"encoding/json"
	"strings"
)

//PowerModule displays a set of button to poweroff, suspend or reboot the computer
type PowerModule struct {
	ModuleBase
	Format, PowerOffText, RebootText, SuspendText string
}

//BuildPower initializes a PowerModule
func BuildPower(ms *ModuleSpec) Module {
	opts := struct {
		Format, PowerOffText, RebootText, SuspendText string
	}{
		Format:       "%P",
		PowerOffText: "\uf011",
		RebootText:   "\uf021",
		SuspendText:  "\uf186",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &PowerModule{
		ModuleBase:   buildModuleBase(ms),
		Format:       opts.Format,
		PowerOffText: opts.PowerOffText,
		RebootText:   opts.RebootText,
		SuspendText:  opts.SuspendText,
	}
}

//Run starts the module
func (m *PowerModule) Run() {
	result := m.Format
	result = strings.Replace(result, "%P", buttonify("poweroff", m.PowerOffText), -1)
	result = strings.Replace(result, "%R", buttonify("reboot", m.RebootText), -1)
	result = strings.Replace(result, "%S", buttonify("systemctl suspend", m.SuspendText), -1)
	output <- m.update(result)
}
