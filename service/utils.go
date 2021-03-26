package service

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
