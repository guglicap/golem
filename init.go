package main

import (
	"encoding/json"
	"log"
)

type Config struct {
	Left   []Module
	Center []Module
	Right  []Module
}

const (
	LEFT   = iota
	CENTER = iota
	RIGHT  = iota
)

func readModules(file []byte) Config {
	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Couldn't decode modules file.", err)
	}
	return config
}

func startModule(bar map[int][]string, m Module, f ModuleHandler) {
	slice := bar[m.position]
	slice = append(slice, "")
	bar[m.position] = slice
	m.index = len(slice) - 1
	go f(m)
}

func spawnModules(config Config) map[int][]string {

	bar := make(map[int][]string)
	startModule(bar, Module{"", 0, 0, LEFT}, padder)

	spawnSlice(bar, config.Left, LEFT)
	spawnSlice(bar, config.Center, CENTER)
	spawnSlice(bar, config.Right, RIGHT)

	startModule(bar, Module{"", 0, 0, RIGHT}, padder)
	return bar
}

func spawnSlice(bar map[int][]string, slice []Module, p int) {
	for _, m := range slice {
		if f, ok := modtypes[m.Type]; ok {
			m.position = p
			startModule(bar, m, f)
		}
	}
}
