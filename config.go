package main

import (
	"encoding/json"

	"github.com/guglicap/golem/modules"
)

type Config struct {
	ErrorColor string //Used when outputting errors to the bar.
	Padding    int    //Space between each module
	Modules    []modules.Module
}

func (c *Config) UnmarshalJSON(data []byte) error {
	var conf struct {
		Padding    int
		ErrorColor string
		Modules    []*modules.ModuleSpec
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		return err
	}
	for _, ms := range conf.Modules {
		mod := buildModule(ms)
		c.Modules = append(c.Modules, mod)
	}
	c.Padding = conf.Padding
	c.ErrorColor = conf.ErrorColor
	return nil
}

func buildModule(ms *modules.ModuleSpec) modules.Module {
	if build, ok := buildFuncs[ms.Handler]; ok {
		return build(ms)
	}
	return modules.BuildUnknownModule(ms)
}

var buildFuncs = map[string]func(ms *modules.ModuleSpec) modules.Module{
	"date":     modules.BuildDate,
	"power":    modules.BuildPower,
	"netinfo":  modules.BuildNet,
	"text":     modules.BuildText,
	"button":   modules.BuildButton,
	"ws":       modules.BuildWs,
	"meminfo":  modules.BuildMem,
	"diskinfo": modules.BuildDisk,
	"launcher": modules.BuildLauncher,
	"whoami":   modules.BuildWhoami,
	"cpu":      modules.BuildCPU,
	"mpd":      modules.BuildMpd,
}
