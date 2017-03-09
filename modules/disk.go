package modules

import (
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/shirou/gopsutil/disk"
)

//DiskModule displays infos about disk usage.
type DiskModule struct {
	ModuleBase
	Mount, Format string
}

//BuildDisk initializes a DiskModule
func BuildDisk(ms *ModuleSpec) Module {
	opts := struct {
		Mount, Format string
	}{
		"/", "%mount, %usePerc",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &DiskModule{
		buildModuleBase(ms),
		opts.Mount,
		opts.Format,
	}
}

//Run starts the module
func (m *DiskModule) Run() {
	for {
		usage, err := disk.Usage(m.Mount)
		if err != nil {
			m.errOutput(err)
			return
		}
		size, used, avail, usedPerc := toGBs(usage.Total), toGBs(usage.Used), toGBs(usage.Free), usage.UsedPercent
		result := m.Format
		result = strings.Replace(result, "%size", strconv.Itoa(size), -1)
		result = strings.Replace(result, "%used", strconv.Itoa(used), -1)
		result = strings.Replace(result, "%avail", strconv.Itoa(avail), -1)
		result = strings.Replace(result, "%usePerc", strconv.FormatFloat(usedPerc, 'f', 1, 64)+"%", -1)
		result = strings.Replace(result, "%mount", usage.Path, -1)
		output <- m.update(result)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
