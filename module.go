package main

import (
	"encoding/json"
	"time"
)

type ModuleHandler func(m Module)
type Duration time.Duration

type Module struct {
	Type     string
	Refresh  Duration
	index    int
	position int
}

type Update struct {
	Position int
	Index    int
	Content  string
}

var modtypes = map[string]ModuleHandler{
	"ws":      ws,
	"syu":     syu,
	"clock":   clock,
	"date":    date,
	"netAddr": netAddr,
	"ping":    ping,
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	temp, err := time.ParseDuration(s)
	if err != nil {
		*d = Duration(0 * time.Second)
	} else {
		*d = Duration(temp)
	}
	return nil
}
