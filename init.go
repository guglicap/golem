package main

import (
	"encoding/json"
	"log"

	"github.com/guglicap/golem/modules"
)

const (
	LEFT   = iota
	CENTER = iota
	RIGHT  = iota
)

//lastPosition holds the position of the latest initialized module.
var lastPosition int

func readConfig(file []byte) Config {
	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Couldn't decode config file.", err)
	}
	return config
}

func startBlock(bar map[int][]string, m modules.Module) {
	pos := m.GetPosition()
	if pos == -1 {
		m.SetPosition(lastPosition)
		pos = lastPosition
	}
	slice := bar[pos]
	slice = append(slice, "")
	bar[pos] = slice
	m.SetIndex(len(slice) - 1)
	lastPosition = pos
	go m.Run()
}

//Initializes the map which holds all of our modules.
func spawnModules(config Config) map[int][]string {
	bar := make(map[int][]string)
	for _, m := range config.Modules {
		startBlock(bar, m)
	}
	return bar
}
