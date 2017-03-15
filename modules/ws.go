package modules

import "encoding/json"

//WsModule is a workspace indicator for bspwm.
type WsModule struct {
	ModuleBase
	WsFocused, WsUnfocused string
	Wm                     string
}

//BuildWs initializes a WsModule
func BuildWs(ms *ModuleSpec) Module {
	opts := struct {
		WsFocused, WsUnfocused string
		Wm                     string
	}{
		WsFocused:   "\uf111",
		WsUnfocused: "\uf10c",
		Wm:          "bspwm",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &WsModule{
		buildModuleBase(ms),
		opts.WsFocused,
		opts.WsUnfocused,
		opts.Wm,
	}
}

//Run starts the module.
func (m *WsModule) Run() {
	switch m.Wm {
	case "bspwm":
		m.bspwm()
	case "i3":
		m.i3()
	}
}
