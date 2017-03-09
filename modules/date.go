package modules

import (
	"encoding/json"
	"time"
)

//DateModule displays the current date/time
type DateModule struct {
	ModuleBase
	Format string
}

//BuildDate initializes a DateModule
func BuildDate(ms *ModuleSpec) Module {
	opts := struct {
		Format string
	}{
		"15:04 02/01/2006",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &DateModule{
		buildModuleBase(ms),
		opts.Format,
	}
}

//Run starts the module
func (m *DateModule) Run() {
	for {
		output <- m.update(time.Now().Format(m.Format))
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
