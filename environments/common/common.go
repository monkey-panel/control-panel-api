package common

import (
	"io"

	. "github.com/monkey-panel/control-panel-api/common/types"
)

type State = uint16

const (
	STATE_BUSY State = iota
	STATE_STOP
	STATE_STOPPING
	STATE_STARTING
	STATE_RUNNING
)

type Environment interface {
	// GetRootPath return the root path of the environment
	GetRootPath() string

	Setup() error
	Start() error
	Stop() error
	Kill() error
	Delete() error

	GetStatus() State
	SetStatus() State
	GetBase() *BaseEnvironment
}

type BaseEnvironment struct {
	Environment

	ID               ID                `json:"id"`                // id of the environment
	Name             string            `json:"name"`              // name of the environment
	Type             string            `json:"type"`              // docker, general
	Command          string            `json:"command"`           // command of the environment
	Env              map[string]string `json:"env"`               // environment variables
	Args             []string          `json:"args"`              // arguments of the command
	RootPath         string            `json:"root"`              // root path of the environment
	WorkingDirectory string            `json:"working_directory"` // working directory of the environment

	ConsoleOutput io.Reader
	SystemOutput  io.Reader
	ConsoleInput  io.Writer

	state State
}

func (e *BaseEnvironment) GetRootPath() string {
	return e.RootPath
}

func (e *BaseEnvironment) SetState(state State) {
	e.state = state
}

func (e *BaseEnvironment) GetStatus() (state State) {
	return e.state
}

func (e *BaseEnvironment) GetBase() *BaseEnvironment {
	return e
}
