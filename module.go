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

var Modtypes = map[string]ModuleHandler{
	"ws":      Ws,
	"syu":     Syu,
	"clock":   Clock,
	"date":    Date,
	"netAddr": NetAddr,
	"ping":    Ping,
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
