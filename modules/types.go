package modules

type ModuleHandler func(m Module)

type Update struct {
	Position int
	Index    int
	Content  string
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
}
