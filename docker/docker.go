package docker

import (
	"context"
	"fmt"
	"github.com/derekhjray/containers/types"
	dtypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
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
		image.Envs = imginfo.Config.Env
	}

	return image, nil
}

func (d *docker) ContainerList(ctx context.Context, option *types.ContainerListOption) ([]*types.Container, error) {
	list, err := d.Client.ContainerList(ctx, dtypes.ContainerListOptions{All: option.All})
	if err != nil {
		return nil, err
	}

	containers := make([]*types.Container, 0, len(list))
	for index := range list {
		container := &types.Container{
			ID:      list[index].ID,
			Image:   list[index].Image,
			ImageID: list[index].ImageID,
			State:   list[index].Status,
			Labels:  list[index].Labels,
			NetMode: list[index].HostConfig.NetworkMode,
			Created: time.Unix(list[index].Created, 0),
		}

		if len(list[index].Names) > 0 {
			container.Name = strings.TrimPrefix(list[index].Names[0], "/")
		}

		containers = append(containers, container)
	}

	return containers, nil
}

func (d *docker) ContainerInspect(ctx context.Context, containerId string) (*types.Container, error) {
	cjson, err := d.Client.ContainerInspect(ctx, containerId)
	if err != nil {
		return nil, err
	}

	container := &types.Container{
		ID:      cjson.ID,
		Name:    strings.TrimPrefix(cjson.Name, "/"),
		ImageID: cjson.Image,
	}

	if container.Created, err = time.Parse(time.RFC3339Nano, cjson.Created); err != nil {
		log.Debugf("Parse container(%s) created time failed, reason: %v", container.Name, err)
	}

	if cjson.State != nil {
		if container.Started, err = time.Parse(time.RFC3339Nano, cjson.State.StartedAt); err != nil {
			log.Debugf("Parse container(%s) started time failed, reason: %v", container.Name, err)
		}

		if container.Finished, err = time.Parse(time.RFC3339Nano, cjson.State.FinishedAt); err != nil {
			log.Debugf("Parse container(%s) finished time failed, reason: %v", container.Name, err)
		}

		container.State = cjson.State.Status
		container.Pid = cjson.State.Pid
	}

	if cjson.Config != nil {
		container.Envs = cjson.Config.Env
		container.Labels = cjson.Config.Labels
		container.User = cjson.Config.User
		if container.User == "" {
			container.User = "root"
		}

		container.Image = cjson.Config.Image
	}

	if cjson.HostConfig != nil {
		container.Privileged = cjson.HostConfig.Privileged
		container.NetMode = string(cjson.HostConfig.NetworkMode)
		container.PidMode = string(cjson.HostConfig.PidMode)
		container.Memory = cjson.HostConfig.Memory
		if cjson.HostConfig.CPUPeriod > 0 {
			container.CPUs = fmt.Sprintf("%0.2f", float64(cjson.HostConfig.CPUQuota)/float64(cjson.HostConfig.CPUPeriod))
		}
	}

	if cjson.NetworkSettings != nil {
		container.IPAddress = cjson.NetworkSettings.IPAddress
		if container.IPAddress == "" {
			for _, settings := range cjson.NetworkSettings.Networks {
				container.IPAddress = settings.IPAddress
				if container.IPAddress != "" {
					break
				}
			}
		}
	}

	return container, nil
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
