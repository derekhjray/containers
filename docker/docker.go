package docker

import (
	"context"
	"github.com/derekhjray/containers/types"
	"github.com/docker/docker/client"
)

func New() (types.Runtime, error) {
	var (
		err error
	)

	runtime := &docker{}
	runtime.Client, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return runtime, nil
}

type docker struct {
	*client.Client
}

func (d *docker) Name() string {
	return types.Docker
}

func (d *docker) ImageList(ctx context.Context) ([]*types.Image, error) {
	return nil, types.ErrMethodNotImplemented
}

func (d *docker) ImageInspect(ctx context.Context, imageId string) (*types.Image, error) {
	imginfo, _, err := d.ImageInspectWithRaw(ctx, imageId)
	if err != nil {
		return nil, err
	}

	image := &types.Image{
		ID:       imginfo.ID,
		RepoTags: imginfo.RepoTags,
	}

	if imginfo.Config != nil {
		image.Labels = imginfo.Config.Labels
		image.Env = imginfo.Config.Env
	}

	return image, nil
}

func (d *docker) ContainerList(ctx context.Context, option *types.ContainerListOption) ([]*types.Container, error) {
	return nil, types.ErrMethodNotImplemented
}

func (d *docker) ContainerInspect(ctx context.Context, containerId string) (*types.Container, error) {
	return nil, types.ErrMethodNotImplemented
}

func (d *docker) ContainerStart(ctx context.Context, containerId string) error {
	return types.ErrMethodNotImplemented
}

func (d *docker) ContainerStop(ctx context.Context, containerId string) error {
	return types.ErrMethodNotImplemented
}

func (d *docker) ContainerPause(ctx context.Context, containerId string) error {
	return types.ErrMethodNotImplemented
}

func (d *docker) ContainerResume(ctx context.Context, containerId string) error {
	return types.ErrMethodNotImplemented
}
