package modules

import "os/exec"

//Init sets global variables found in the config needed by the modules and returns a channel modules use to communicate with the main goroutine.
func Init(errColor string) chan Update {
	errorColor = errColor
	output = make(chan Update)
	return output
}

//inPATH searches for a program in $PATH, if not found returns an error message.

func inPATH(program string) string {
	_, err := exec.LookPath(program)
	if err != nil {
		return colorize(errorColor, "Can't find "+program+" in $PATH")
	}
	return ""
}

func colorize(color, s string) string {
	return "%{F" + color + "}" + s + "%{F-}"
}

func toGBs(blks uint64) int {
	return int(blks) / 1024 / 1024 / 1024
}

func toMBs(blks uint64) int {
	return int(blks) / 1024 / 1024
}

//NOTE: in order for this to work you need to pipe the output of lemonbar to sh
func buttonify(command, s string) string {
	return "%{A:" + command + ":}" + s + "%{A}"
}

func errOutput(m Module, err error) {
	output <- Update{m.slot, m.colors, err.Error()}
}
