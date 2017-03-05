package modules

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/mem"
)

type MemModule struct {
	ModuleBase
	Format string
}

func BuildMem(ms *ModuleSpec) Module {
	opts := struct {
		Format string
	}{
		"%usedMB / %totalMB",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &MemModule{
		buildModuleBase(ms),
		opts.Format,
	}
}

func (m *MemModule) Run() {
	for {
		vm, err := mem.VirtualMemory()
		if err != nil {
			m.errOutput(err)
			return
		}
		free, total, avail, used, usedPerc, cached, shared := toMBs(vm.Free), toMBs(vm.Total), toMBs(vm.Available), toMBs(vm.Used), vm.UsedPercent, toMBs(vm.Cached), toMBs(vm.Shared)
		result := m.Format
		result = strings.Replace(result, "%free", strconv.Itoa(free), -1)
		result = strings.Replace(result, "%total", strconv.Itoa(total), -1)
		result = strings.Replace(result, "%avail", strconv.Itoa(avail), -1)
		result = strings.Replace(result, "%used", strconv.Itoa(used), -1)
		result = strings.Replace(result, "%usePerc", strconv.FormatFloat(usedPerc, 'f', 1, 64), -1)
		result = strings.Replace(result, "%cached", strconv.Itoa(cached), -1)
		result = strings.Replace(result, "%shared", strconv.Itoa(shared), -1)
		output <- Update{m.slot, m.colors, result}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
