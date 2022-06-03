package cmd

import (
	"os"
	"os/exec"
	"runtime"
)

const windowOS = "windows"

func ClearScreen() {
	system := runtime.GOOS
	if system == windowOS {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
