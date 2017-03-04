package modules

import (
	"fmt"
	"os/user"
	"strings"
	"time"

	"strconv"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func powerHub(m Module) {
	result := m.options.PowerHubFormat
	result = strings.Replace(result, "%P", buttonify("poweroff", m.options.PowerOffText), -1)
	result = strings.Replace(result, "%R", buttonify("reboot", m.options.RebootText), -1)
	result = strings.Replace(result, "%S", buttonify("systemctl suspend", m.options.SuspendText), -1)
	fmt.Println(m.options)
	output <- Update{m.slot, m.colors, result}
}

func whoami(m Module) {
	u, err := user.Current()
	if err != nil {
		errOutput(m, err)
		return
	}
	result := m.options.WhoamiFormat
	result = strings.Replace(result, "%uname", u.Username, -1)
	result = strings.Replace(result, "%name", u.Name, -1)
	result = strings.Replace(result, "%gid", u.Gid, -1)
	result = strings.Replace(result, "%uid", u.Uid, -1)
	output <- Update{m.slot, m.colors, result}
}

func diskinfo(m Module) {
	for {
		usage, err := disk.Usage(m.options.DiskMount)
		if err != nil {
			errOutput(m, err)
			return
		}
		size, used, avail, usedPerc := toGBs(usage.Total), toGBs(usage.Used), toGBs(usage.Free), usage.UsedPercent
		result := m.options.DiskInfoFormat
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

func meminfo(m Module) {
	for {
		vm, err := mem.VirtualMemory()
		if err != nil {
			errOutput(m, err)
			return
		}
		free, total, avail, used, usedPerc, cached, shared := toMBs(vm.Free), toMBs(vm.Total), toMBs(vm.Available), toMBs(vm.Used), vm.UsedPercent, toMBs(vm.Cached), toMBs(vm.Shared)
		result := m.options.MemInfoFormat
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
