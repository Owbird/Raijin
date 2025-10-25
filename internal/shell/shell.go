package shell

import "os/exec"

type ShellCmd struct {
	Cmd  string
	Dir  string
	Args []string
}

func Run(shellCmd ShellCmd) ([]byte, error) {
	pendingCmd := exec.Command(shellCmd.Cmd, shellCmd.Args...)
	pendingCmd.Dir = shellCmd.Dir

	println(pendingCmd.String())

	return pendingCmd.Output()
}
