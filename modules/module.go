package modules

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type ModuleHandler func(m Module)

type Update struct {
	Position int
	Index    int
	Content  string
}

var modtypes = map[string]ModuleHandler{
	"ws":      ws,
	"syu":     syu,
	"date":    date,
	"netAddr": netAddr,
	"ping":    ping,
	"pad":     padder,
}

var (
	output     chan Update
	errorColor string
)

func Init(errColor string) chan Update {
	errorColor = errColor
	output = make(chan Update)
	getDefaultInterface()
	return output
}

type Module struct {
	handler  ModuleHandler
	Position int
	Index    int
	runOnce  bool
	refresh  time.Duration
	options  *Options
}

func (m Module) Run() {
	go m.handler(m)
}

func (m *Module) UnmarshalJSON(data []byte) error {
	var temp struct {
		Handler  string
		Position string
		Refresh  string
		Opts     *Options `json:"Options"`
	}
	log.Println(temp.Opts, temp)
	var module Module
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	if f, ok := modtypes[temp.Handler]; ok {
		module.handler = f
	} else {
		module.handler = func(m Module) {
			output <- Update{m.Position, m.Index, "Can't find module type " + temp.Handler}
		}
		return nil
	}
	switch strings.ToLower(temp.Position) {
	case "left":
		module.Position = 0
	case "center":
		module.Position = 1
	case "right":
		module.Position = 2
	default:
		module.Position = 0
	}
	dur, err := time.ParseDuration(temp.Refresh)
	if err != nil {
		module.runOnce = true
	} else {
		module.refresh = dur
		module.runOnce = false
	}
	if temp.Opts == nil {
		module.options = defaultOptions
	} else {
		module.options = temp.Opts
	}
	*m = module
	return nil
}
