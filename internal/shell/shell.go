package shell

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

type ShellCmd struct {
	Cmd  string
	Dir  string
	Args []string
}

func Run(shellCmd ShellCmd) ([]byte, error) {
	pendingCmd := exec.Command(shellCmd.Cmd, shellCmd.Args...)
	pendingCmd.Dir = shellCmd.Dir

	output := &bytes.Buffer{}
	pendingCmd.Stdout = io.MultiWriter(os.Stdout, output)
	pendingCmd.Stderr = os.Stderr

	println(pendingCmd.String())
	err := pendingCmd.Run()
	return output.Bytes(), err
}
