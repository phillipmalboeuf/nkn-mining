package container

import "os/exec"

type IContainer interface {
	InitEnvironment(path string, workPath string) bool
	AsyncRun(args []string, input string) (cmd *exec.Cmd, err error)
	SyncRun(args []string, input string) (output string, err error)
	Stop()
	Status() int
}
