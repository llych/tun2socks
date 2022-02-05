package command

import (
	"os/exec"
)

func Run(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	data, err := cmd.CombinedOutput()
	return string(data), err
}
