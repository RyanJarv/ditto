package main

import (
	"errors"
	"github.com/ryanjarv/ditto/pkg"
	"os"
	"os/exec"
)

func main() {
	err := bitm.Run(os.Args)
	if errors.Is(err, &exec.ExitError{}) {
		os.Exit(err.(*exec.ExitError).ExitCode())
	} else if err != nil {
		os.Exit(234)
	} else {
		os.Exit(0)
	}
}
