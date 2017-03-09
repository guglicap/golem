package modules

import (
	"encoding/json"
	"errors"
	"strings"
)

//UnknownModule we use this when there's something wrong in the config.
type UnknownModule struct {
	ModuleBase
	ModuleName string
}

//BuildUnknownModule initializes an UnknownModule
func BuildUnknownModule(ms *ModuleSpec) Module {
	return &UnknownModule{
		buildModuleBase(ms),
		ms.Handler,
	}
}

//Run starts the module
func (m *UnknownModule) Run() {
	m.errOutput(errors.New("Unknown module " + m.ModuleName))
}

//TextModule displays text.
type TextModule struct {
	ModuleBase
	Text string
}

//BuildText initializes a TextModule
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

//Run starts the module
func (m *TextModule) Run() {
	output <- m.update(m.Text)
}

//ButtonModule displays a button
type ButtonModule struct {
	ModuleBase
	Command, Text string
}

//BuildButton initializes a ButtonModule
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

//Run starts the module
func (m *ButtonModule) Run() {
	cmd, txt := strings.Fields(m.Command), m.Text //We get the fields to separate the command from the args.
	if len(cmd) == 0 {
		m.errOutput(errors.New("You didn't specify a command."))
		return
	}
	if len(txt) == 0 {
		txt = m.Command
	}
	if ok := inPATH(cmd[0]); !ok {
		m.errOutput(errors.New("Can't find " + cmd[0] + " in $PATH"))
		return
	}
	output <- m.update(buttonify(m.Command, txt))
}
