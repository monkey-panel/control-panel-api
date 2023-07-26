//go:build windows
// +build windows

package general

import (
	"io"
	"os"
	"os/exec"
	"path"
)

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

	process.StderrPipe()

	go e.handleClose()
	return
}

func (e *General) handleClose() (err error) {
	if err = e.process.Wait(); err != nil {
		return
	}

	return
}
