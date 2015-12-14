package exec

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Cmd wraps the command execution pattern required in os/exec package.
// The command is executed with the supplied arguments and the output is returned
// to the caller. If an error occurs during the command execution, an error is
// returned with the messages read from the stderr of the executed command.
func Cmd(executable string, args ...string) (string, error) {
	cmd := exec.Command(executable, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	outBuf, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	out := string(outBuf)

	errBuf, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println(errBuf)
		return "", err
	}
	errMsg := string(errBuf)

	if err := cmd.Wait(); err != nil {
		if errMsg != "" {
			return out, errors.New(errMsg)
		}

		return out, err
	}

	return out, nil
}
