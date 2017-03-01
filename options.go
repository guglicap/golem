package main

import (
	"encoding/json"
	"log"
)

type Options struct {
	Padding     int
	ErrorColor  string
	PingAddr    string
	WsFocused   string
	WsUnfocused string
	ClockFormat string
	DateFormat  string
}

func defaultOptions() Options {
	return Options{
		8,
		"#c98d2c",
		"google.com",
		"\uf111",
		"\uf10c",
		"15:04",
		"2/01/2006",
	}
}

func loadOptions(file []byte) Options {
	opt := defaultOptions()
	err := json.Unmarshal(file, &opt)
	if err != nil {
		log.Fatal("Couldn't decode options file.", err)
	}
	return opt
}
