package general

import (
	"os/exec"

	"github.com/monkey-panel/control-panel-api/environments/common"
)

type General struct {
	common.BaseEnvironment
	process *exec.Cmd
}

func (e *General) Setup() (err error) {
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
