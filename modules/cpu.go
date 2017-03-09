package modules

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//CPUModule displays CPU Usage %
type CPUModule struct {
	ModuleBase
	Format string
	PerCPU bool
}

//This is per-core actually.
type lastCPUStat struct {
	busy  float64
	total float64
}

//BuildCPU initializes a CPUModule
func BuildCPU(ms *ModuleSpec) Module {
	opts := struct {
		Format string
		PerCPU bool
	}{
		"%core: %usage%  ",
		true,
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &CPUModule{
		buildModuleBase(ms),
		opts.Format,
		opts.PerCPU,
	}
}

//Run starts the module
func (m *CPUModule) Run() {
	cpus := make([][]string, 0)         //Hold fields of lines of /proc/stats,
	lastStats := make([]lastCPUStat, 0) //Holds the last stats for the cores
	file, err := os.Open("/proc/stat")
	if err != nil {
		m.errOutput(err)
		return
	}
	defer file.Close()
	for {
		i := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { //Read our file
			line := scanner.Text()
			if strings.HasPrefix(line, "cpu") { //We're only interested in those lines.
				if len(cpus) <= i {
					cpus = append(cpus, strings.Fields(line))
					lastStats = append(lastStats, lastCPUStat{
						busy:  0,
						total: 0,
					})
				} else {
					cpus[i] = strings.Fields(line)
				}
				i++
			}
		}
		if err := scanner.Err(); err != nil { //On EOF err will be nil
			m.errOutput(err)
			return
		}
		_, err = file.Seek(0, 0) //Rewind our file, so next read the kernel will get us new stats
		if err != nil {
			m.errOutput(err)
			return
		}
		var result string
		for i, c := range cpus { //Here comes the bad code
			if i == 0 && m.PerCPU { //The first line of the /proc/stat includes all of the cores.
				continue
			}
			var core string
			var busy, total float64
			for j, s := range c {
				if j == 0 { //The first field in the line is the cpu id, such as "cpu0"
					core = s
					continue
				}
				f, err := strconv.ParseFloat(s, 64) //Get our values as float
				if err != nil {
					log.Println("You gotta be kidding. Failed to parse", s)
					continue
				}
				total += f
				if j != 4 { //5th field of the line is the idle time
					busy += f
				}
			}
			usage := (busy - lastStats[i].busy) / (total - lastStats[i].total) * 100
			if busy <= lastStats[i].busy { //Copied from gopsutils, I have no idea how it works
				usage = 0
			}
			if total <= lastStats[i].total { //Same
				usage = 1
			}
			//log.Printf("Busy: %.2f, Idle: %.2f, Usage: %s\n", busy, idle, usage)
			result += strings.Replace(strings.Replace(m.Format, "%core", core, -1), "%usage", strconv.FormatFloat(usage, 'f', 2, 64), -1)
			//log.Println(result)
			lastStats[i].busy, lastStats[i].total = busy, total
			if !m.PerCPU { //If we didn't skip the first iteration, this has to be true, so leave this loop
				break
			}
		}
		//And we're done
		output <- m.update(result)
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
