package bash

import (
	"os"
	"os/exec"
)

func Execute(cmds ...string) {

	for _, cmd := range cmds {
		command := exec.Command("bash", "-c", cmd)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Run()
	}
}
