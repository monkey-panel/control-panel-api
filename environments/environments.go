package environments

import "github.com/monkey-panel/control-panel-api/environments/common"

type Type = string

const (
	TYPE_DOCKER  Type = "docker"
	TYPE_GENERAL Type = "general"
)

// var environments = map[Type]common.BaseEnvironment{
// 	// TYPE_DOCKER: docker.Docker{},
// }

func init() {
}

func Create(env common.BaseEnvironment) (err error) {
	return
}
