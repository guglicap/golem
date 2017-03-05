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

type Module interface {
	Run()
	SetPosition(int)
	GetPosition() int
	SetIndex(int)
	GetIndex() int
}

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

type ModuleBase struct {
	slot    Slot
	refresh time.Duration //How often modules refresh. Note that not every module does.
	runOnce bool          //When true modules that would normally refresh exit after one iteration.
	colors  Colors
}

func (mb *ModuleBase) GetPosition() int {
	return mb.slot.Position
}
func (mb *ModuleBase) SetPosition(i int) {
	mb.slot.Position = i
}
func (mb *ModuleBase) GetIndex() int {
	return mb.slot.Index
}
func (mb *ModuleBase) SetIndex(i int) {
	mb.slot.Index = i
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
		mb.slot.Position = -1
	}
	mb.colors = ms.Colors
	dur, err := time.ParseDuration(ms.Refresh)
	if err != nil || dur < 1*time.Second {
		mb.runOnce = true
	} else {
		mb.refresh = dur
		mb.runOnce = false
	}
	return mb
}
