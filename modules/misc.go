package modules

import (
	"encoding/json"
	"errors"
)

type UnknownModule struct {
	ModuleBase
	ModuleName string
}

func BuildUnknownModule(ms *ModuleSpec) Module {
	return &UnknownModule{
		buildModuleBase(ms),
		ms.Handler,
	}
}

func (m *UnknownModule) Run() {
	output <- Update{m.slot, m.colors, "Unknown module " + m.ModuleName}
}

type TextModule struct {
	ModuleBase
	Text string
}

func BuildText(ms *ModuleSpec) Module {
	var opts struct {
		Text string
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &TextModule{
		buildModuleBase(ms),
		opts.Text,
	}
}

func (m *TextModule) Run() {
	output <- Update{m.slot, m.colors, m.Text}
}

type ButtonModule struct {
	ModuleBase
	Command, Text string
}

func BuildButton(ms *ModuleSpec) Module {
	var opts struct {
		Command, Text string
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &ButtonModule{
		ModuleBase: buildModuleBase(ms),
		Command:    opts.Command,
		Text:       opts.Text,
	}
}

func (m *ButtonModule) Run() {
	cmd, txt := m.Command, m.Text
	if len(cmd) == 0 {
		m.errOutput(errors.New("You didn't specify a command."))
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
