package containers

import (
	"github.com/derekhjray/containers/docker"
	"github.com/derekhjray/containers/types"
)

func NewRuntime() (types.Runtime, error) {
	return docker.New()
}
