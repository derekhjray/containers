package types

import (
	"context"
	"errors"
)

const (
	Docker     = "docker"
	Containerd = "containerd"
	CriO       = "cri-o"
	Podman     = "podman"
)

type Options struct {
	Host string
}

var (
	ErrMethodNotImplemented = errors.New("method not implemented")
)

type ContainerListOption struct {
	All bool
}

type Runtime interface {
	Name() string
	ImageList(ctx context.Context) ([]*Image, error)
	ImageInspect(ctx context.Context, imageId string) (*Image, error)
	ContainerList(ctx context.Context, option *ContainerListOption) ([]*Container, error)
	ContainerInspect(ctx context.Context, containerId string) (*Container, error)
	ContainerStart(ctx context.Context, containerId string) error
	ContainerStop(ctx context.Context, containerId string) error
	ContainerPause(ctx context.Context, containerId string) error
	ContainerResume(ctx context.Context, containerId string) error
}
