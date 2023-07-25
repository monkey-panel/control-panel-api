package general

import (
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/monkey-panel/control-panel-api/environments"
)

type General struct {
	environments.BaseEnvironment
	process *exec.Cmd
}

func (e *General) Setup() (err error) {
	return
}

func (e *General) Start() (err error) {
	process := exec.Command(e.Command, e.Args...)
	process.Dir = path.Join(e.GetRootPath(), e.WorkingDirectory)

	process.Env = append(process.Env, os.Environ()...)
	process.Env = append(process.Env, "HOME="+e.GetRootPath())
	for k, v := range e.Env {
		process.Env = append(process.Env, k+"="+v)
	}

	stdout, err := process.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := process.StderrPipe()
	if err != nil {
		return err
	}
	stdin, err := process.StdinPipe()
	if err != nil {
		return err
	}
	e.ConsoleInput = stdin
	e.ConsoleOutput = io.MultiReader(stdout, stderr)
	e.process = process

	if err = process.Start(); err != nil {
		return
	}

	return
}

func (e *General) Stop() (err error) {
	return
}

func (e *General) Kill() (err error) {
	return
}

func (e *General) Delete() (err error) {
	return
}
