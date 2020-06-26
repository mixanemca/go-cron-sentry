package local

import (
	"bytes"
	"io/ioutil"
	"os/exec"
)

// Run execute commands with args
func (c *client) Run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e

	if c.quite {
		cmd.Stdout = ioutil.Discard
		cmd.Stderr = ioutil.Discard
	}

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return o.String() + e.String(), nil
}
