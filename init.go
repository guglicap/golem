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

type Config struct {
	ErrorColor string
	Padding    int
	Modules    []modules.Module
}

func readModules(file []byte) Config {
	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Couldn't decode modules file.", err)
	}
	return config
}

func startModule(bar map[int][]string, m modules.Module) {
	slice := bar[m.Position]
	slice = append(slice, "")
	bar[m.Position] = slice
	m.Index = len(slice) - 1
	m.Run()
}

func spawnModules(config Config) map[int][]string {

	bar := make(map[int][]string)
	for _, m := range config.Modules {
		startModule(bar, m)
	}
	return bar
}
