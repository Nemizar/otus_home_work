package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for k, v := range env {
		if v.NeedRemove {
			err := os.Unsetenv(k)
			if err != nil {
				return
			}

			continue
		}

		err := os.Setenv(k, v.Value)
		if err != nil {
			return
		}
	}

	childCmd := exec.Command(cmd[0], cmd[1:]...) //nolint: gosec
	childCmd.Env = os.Environ()
	childCmd.Stdin = os.Stdin
	childCmd.Stdout = os.Stdout
	childCmd.Stderr = os.Stderr

	if err := childCmd.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}

	return
}
