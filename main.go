package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	out     chan Update
	options Options
	tab     string
)

const (
	DEBUG = true
)

func main() {
	if !DEBUG {
		logFile, err := os.Create("/home/guglielmo/.bin/desktop/gobar.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	if len(os.Args) < 3 {
		log.Fatal("Usage: golem <config_file> <options_file>")
	}
	configFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	optionsFile, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	config := readModules(configFile)
	options = loadOptions(optionsFile)
	getDefaultInterface()
	log.Println(defaultNetInterface)
	tab = strings.Repeat(" ", options.Padding)
	out = make(chan Update)
	bar := spawnModules(config)

	for m := range out {
		bar[m.Position][m.Index] = m.Content
		for _, k := range [3]int{LEFT, CENTER, RIGHT} {
			switch k {
			case LEFT:
				fmt.Print("%{l}")
			case CENTER:
				fmt.Print("%{c}")
			case RIGHT:
				fmt.Print("%{r}")
			}
			fmt.Print(strings.Join(bar[k], tab))
		}
		fmt.Println()
	}
}
