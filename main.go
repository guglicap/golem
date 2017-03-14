package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"flag"

	"github.com/guglicap/golem/modules"
)

const (
	DEBUG = false //When true logOutput is set to stdout.
)

func setColors(u modules.Update) string {
	result := u.Content
	if len(u.Color.Background) != 0 {
		result = "%{B" + u.Color.Background + "}" + result + "%{B-}"
	}
	if len(u.Color.Foreground) != 0 {
		result = "%{F" + u.Color.Foreground + "}" + result + "%{F-}"
	}
	return result
}

var configfl = flag.String("config", "golem.json", "config file")

func main() {
	flag.Parse()
	if !DEBUG {
		logFile, err := os.Create("golem.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	configFile, err := ioutil.ReadFile(*configfl)
	if err != nil {
		log.Fatal(err)
	}

	bar := spawnModules(readConfig(configFile))
	//Reads Updates from the channel

	for u := range out {
		//Sets the corresponding bar "slot" to containt the update we just received
		bar[u.Slot.Position][u.Slot.Index] = setColors(u)
		for _, k := range [3]int{LEFT, CENTER, RIGHT} { //"Flushes" the array to lemonbar.
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
