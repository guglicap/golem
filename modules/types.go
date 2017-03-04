package modules

type ModuleHandler func(m Module)

type Update struct {
	Slot    Slot
	Color   Colors
	Content string
}

type Colors struct {
	Background string
	Foreground string
}

type Slot struct {
	Position int
	Index    int
}

var modtypes = map[string]ModuleHandler{
	"workspaces": ws,
	"date":       date,
	"netAddress": netAddress,
	"text":       text,
	"powertray":  powerHub,
	"button":     button,
	"whoami":     whoami,
	"icontray":   icontray,
	"diskinfo":   diskinfo,
	"meminfo":    meminfo,
}
