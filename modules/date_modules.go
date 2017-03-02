package modules

import "time"

func date(m Module) {
	for {
		output <- Update{m.Position, m.Index, time.Now().Format(m.options.DateFormat)}
		if m.runOnce {
			return
		}
		time.Sleep(m.refresh)
	}
}
