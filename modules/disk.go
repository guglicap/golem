package modules

import (
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/shirou/gopsutil/disk"
)

type DiskModule struct {
	ModuleBase
	DiskMount, Format string
}

func BuildDisk(ms *ModuleSpec) Module {
	opts := struct {
		DiskMount, Format string
	}{
		"/", "%mount, %usePerc",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &DiskModule{
		buildModuleBase(ms),
		opts.DiskMount,
		opts.Format,
	}
}

func (m *DiskModule) Run() {
	for {
		usage, err := disk.Usage(m.DiskMount)
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
		output <- Update{m.slot, m.colors, result}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
