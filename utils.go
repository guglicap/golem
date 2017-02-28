package main

import "os/exec"
import "time"

func inPATH(program string) string {
	_, err := exec.LookPath("pacman")
	if err != nil {
		return errorColor("-Syu: Can't find " + program + " in $PATH")
	}
	return ""
}

func errorColor(s string) string {
	return "%{F#c98d2c}" + s + "%{F-}"
}

func toTime(d Duration) time.Duration {
	return time.Duration(d)
}

func checkDuration(d Duration) bool {
	return time.Duration(d) == 0*time.Second
}
