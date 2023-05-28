package helpers

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	// cmd := exec.Command("cls") // for windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
