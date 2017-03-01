package main

import "time"

func Clock(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		out <- Update{m.position, m.index, time.Now().Format(options.ClockFormat)}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}

func Date(m Module) {
	runOnce := checkDuration(m.Refresh)
	for {
		out <- Update{m.position, m.index, time.Now().Format(options.DateFormat)}
		if runOnce {
			return
		}
		time.Sleep(toTime(m.Refresh))
	}
}
