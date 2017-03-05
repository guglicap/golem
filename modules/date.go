package modules

import (
	"encoding/json"
	"time"
)

type DataModule struct {
	ModuleBase
	Format string
}

func BuildDate(ms *ModuleSpec) Module {
	opts := struct {
		Format string
	}{
		"15:04 02/01/2006",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &DataModule{
		buildModuleBase(ms),
		opts.Format,
	}
}

func (m *DataModule) Run() {
	for {
		output <- Update{m.slot, m.colors, time.Now().Format(m.Format)}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
