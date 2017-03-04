package modules

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	output     chan Update //Channel used to communicate with the main goroutine. Same for all the modules.
	errorColor string
)

//Module corresponds to a "slot" in the bar.
type Module struct {
	handler ModuleHandler //Function called by this module on Run. I'm not very proud of how I'm doing this, but hey, it works.
	slot    Slot
	refresh time.Duration //How often modules refresh. Note that not every module does.
	runOnce bool          //When true modules that would normally refresh exit after one iteration.
	colors  Colors
	options *Options //Also not very proud of how I'm doing this. Holds the options for this module, set to defaultOptions when there are none.
}

//Run starts the module. I'm bad at comments.
func (m Module) Run() {
	go m.handler(m)
}

//UnmarshalJSON is a custom JSON Unmarshaler for the Module struct
func (m *Module) UnmarshalJSON(data []byte) error {
	var temp struct {
		Handler  string
		Position string
		Refresh  string
		Colors   Colors
		options  *Options
	}
	var module Module
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	if f, ok := modtypes[temp.Handler]; ok {
		module.handler = f
	} else {
		module.handler = func(m Module) {
			errOutput(m, errors.New("Can't find module "+temp.Handler))
		}
	}
	if defOpt, ok := defaultOptions[temp.Handler]; ok {
		var opts struct {
			Opts *Options `json:"Options"`
		}
		opts.Opts = new(Options)
		*(opts.Opts) = *defOpt
		json.Unmarshal(data, &opts)
		module.options = opts.Opts
	}
	pos := &module.slot.Position
	switch strings.ToLower(temp.Position) {
	case "left":
		*pos = 0
	case "center":
		*pos = 1
	case "right":
		*pos = 2
	default:
		*pos = -1 //This means we didn't have a Position property in the config and we'll use the last one.
	}
	module.colors = temp.Colors
	dur, err := time.ParseDuration(temp.Refresh)
	if err != nil || dur < 1*time.Second {
		module.runOnce = true
	} else {
		module.refresh = dur
		module.runOnce = false
	}
	*m = module
	return nil
}

func (m *Module) SetPosition(p int) {
	m.slot.Position = p
}

func (m *Module) SetIndex(i int) {
	m.slot.Index = i
}

func (m *Module) GetIndex() int {
	return m.slot.Index
}

func (m *Module) GetPosition() int {
	return m.slot.Position
}
