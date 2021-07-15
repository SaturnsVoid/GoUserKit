package core

import (
	"os/exec"
)

func run(cmd string) {
	exec.Command("cmd", "/C", cmd).Run()
}
