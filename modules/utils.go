package modules

import "os/exec"

func inPATH(program string) string {
	_, err := exec.LookPath(program)
	if err != nil {
		return colorize(errorColor, "-Syu: Can't find "+program+" in $PATH")
	}
	return ""
}

func colorize(color, s string) string {
	return "%{F" + color + "}" + s + "%{F-}"
}
