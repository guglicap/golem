package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/guglicap/golem/modules"
)

const (
	DEBUG = true
)

func main() {
	if !DEBUG {
		logFile, err := os.Create("golem.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	if len(os.Args) < 2 {
		log.Fatal("Usage: golem <config_file>")
	}
	configFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	config := readModules(configFile)
	tab := strings.Repeat(" ", config.Padding)
	out := modules.Init(config.ErrorColor)

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
