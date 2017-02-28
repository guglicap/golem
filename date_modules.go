package main

import "time"

func clock(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		out <- Update{m.position, m.index, time.Now().Format("15:04")}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}

func date(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		out <- Update{m.position, m.index, time.Now().Format("2/01/2006")}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}
