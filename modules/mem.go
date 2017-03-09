package modules

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"bufio"
	"os"
	"time"
)

//MemModule displays info about memory usage.
type MemModule struct {
	ModuleBase
	Format string
}

//BuildMem initializes a MemModule
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

//Run starts the module.
func (m *MemModule) Run() {
	mem, err := os.Open("/proc/meminfo")
	defer mem.Close()
	meminfo := make([]string, 3)
	if err != nil {
		m.errOutput(err)
		return
	}
	for {
		scanner := bufio.NewScanner(mem) //Same story as CPUModule
		i := 0
		for scanner.Scan() {
			if i <= 2 {
				if f := strings.Fields(scanner.Text()); len(f) > 1 { //This check is here just in case.
					meminfo[i] = f[1][:len(f[1])-3] //Divide by 1000, the kernel gives us the values in kB, we want MB
				}
				i++
			} else {
				break
			}
		}
		_, err := mem.Seek(0, 0)
		if err != nil {
			m.errOutput(err)
			return
		}
		total, err := strconv.Atoi(meminfo[0])
		if err != nil {
			log.Println(err)
			total = 1
		}
		avail, err := strconv.Atoi(meminfo[2])
		if err != nil {
			log.Println(err)
			avail = 1
		}
		used, usedPerc := total-avail, float64(total-avail)/float64(total)*100
		var result string
		result = strings.Replace(m.Format, "%total", meminfo[0][:], -1)
		result = strings.Replace(result, "%free", meminfo[1], -1)
		result = strings.Replace(result, "%avail", meminfo[2], -1)
		result = strings.Replace(result, "%used", strconv.Itoa(used), -1)
		result = strings.Replace(result, "%usePerc", strconv.FormatFloat(usedPerc, 'f', 2, 64), -1)
		output <- m.update(result)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
