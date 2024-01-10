package main

import (
	"os/exec"
)

// step struct
type step struct {
	name    string
	exe     string
	args    []string
	message string
	proj    string
}

// newStep
func newStep(name, exe, message, proj string, args []string) step {
	return step{
		name:    name,
		exe:     exe,
		message: message,
		args:    args,
		proj:    proj,
	}
}

// execute
func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj
	// fmt.Println("cmd.Dir", s.args, cmd.Dir)

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}
	return s.message, nil
}
