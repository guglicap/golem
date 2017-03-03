package modules

import (
	"encoding/json"
	"strings"
	"time"
)

var (
	output     chan Update //Channel used to communicate with the main goroutine. Same for all the modules.
	errorColor string
)

//Module corresponds to a "slot" in the bar.
type Module struct {
	handler  ModuleHandler //Function called by this module on Run. I'm not very proud of how I'm doing this, but hey, it works.
	position int           //LEFT = 0, CENTER = 1, RIGHT = 2
	index    int
	refresh  time.Duration //How often modules refresh. Note that not every module does.
	runOnce  bool          //When true modules that would normally refresh exit after one iteration.
	options  *Options      //Also not very proud of how I'm doing this. Holds the options for this module, set to defaultOptions when there are none.
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
			output <- Update{m.position, m.index, colorize(errorColor, "Can't find module "+temp.Handler)}
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
	switch strings.ToLower(temp.Position) {
	case "left":
		module.position = 0
	case "center":
		module.position = 1
	case "right":
		module.position = 2
	default:
		module.position = -1 //This means we didn't have a Position property in the config and we'll use the last one.
	}
	dur, err := time.ParseDuration(temp.Refresh)
	if err != nil {
		module.runOnce = true
	} else {
		module.refresh = dur
		module.runOnce = false
	}
	*m = module
	return nil
}

func (m *Module) SetPosition(p int) {
	m.position = p
}

func (m *Module) SetIndex(i int) {
	m.index = i
}

func (m *Module) GetPosition() int {
	return m.position
}
