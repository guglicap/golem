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

//Module includes methods common to all modules.
type Module interface {
	Run()
	SetPosition(int)
	GetPosition() int
	SetIndex(int)
	GetIndex() int
}

//ModuleSpec is what we unmarshal the config encoding into.
type ModuleSpec struct {
	Handler  string
	Position string
	Refresh  string
	Colors   struct {
		Background string
		Foreground string
	}
	Options json.RawMessage
}

//ModuleBase holds fields common to all modules
type ModuleBase struct {
	slot    Slot
	refresh time.Duration //How often modules refresh. Note that not every module does.
	runOnce bool          //When true modules that would normally refresh exit after one iteration.
	colors  Colors
}

//GetPosition returns the module Position
func (mb *ModuleBase) GetPosition() int {
	return mb.slot.Position
}

//SetPosition sets the module Position
func (mb *ModuleBase) SetPosition(i int) {
	mb.slot.Position = i
}

//GetIndex returns the module Index
func (mb *ModuleBase) GetIndex() int {
	return mb.slot.Index
}

//SetIndex sets the module Index
func (mb *ModuleBase) SetIndex(i int) {
	mb.slot.Index = i
}

func (mb *ModuleBase) update(content string) Update {
	return Update{mb.slot, mb.colors, content}
}

func buildModuleBase(ms *ModuleSpec) ModuleBase {
	var mb ModuleBase
	switch strings.ToLower(ms.Position) {
	case "left":
		mb.slot.Position = 0
	case "center":
		mb.slot.Position = 1
	case "right":
		mb.slot.Position = 2
	default:
		mb.slot.Position = -1 //This tells the code in the main package to use the same position as the last module initialized
	}
	mb.colors = ms.Colors
	dur, err := time.ParseDuration(ms.Refresh)
	if err != nil || dur < 1*time.Second { //Minimum refresh time allowed is 1 sec
		mb.runOnce = true
	} else {
		mb.refresh = dur
		mb.runOnce = false
	}
	return mb
}
