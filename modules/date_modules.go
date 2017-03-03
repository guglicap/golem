package modules

import "time"

func date(m Module) {
	for {
		output <- Update{m.position, m.index, time.Now().Format(m.options.DateFormat)}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
