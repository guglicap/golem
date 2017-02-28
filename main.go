package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	out chan Update
)

func main() {
	if len(os.Args) != 3 {
		logFile, err := os.Create("/home/guglielmo/.bin/desktop/gobar.log")
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
	bar := spawnModules(config)
	out = make(chan Update)
	getDefaultInterface()
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
