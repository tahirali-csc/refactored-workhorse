package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/workhorse/worker/pkg/engine"
	"io"
	"io/ioutil"
)

type dockerEngine struct {
	dockerClient client.APIClient
}

func NewDockerEngine() (engine.Engine, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	engine := &dockerEngine{dockerClient: cli}
	return engine, err
}

func (de *dockerEngine) Start(ctx context.Context, spec *engine.Spec, step *engine.Step) error {
	return de.dockerClient.ContainerStart(ctx, step.Metadata.UID, types.ContainerStartOptions{})
}

func (de *dockerEngine) Create(ctx context.Context, spec *engine.Spec, step *engine.Step) error {

	reader, err := de.dockerClient.ImagePull(ctx, step.Docker.Image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	io.Copy(ioutil.Discard, reader)
	reader.Close()

	_, err = de.dockerClient.ContainerCreate(ctx,
		toConfig(spec, step),
		toHostConfig(spec, step), nil, step.Metadata.UID)

	return err
}

func (de *dockerEngine) Tail(ctx context.Context, spec *engine.Spec, step *engine.Step) (io.ReadCloser, error) {
	opts := types.ContainerLogsOptions{
		Follow:     true,
		ShowStdout: true,
		ShowStderr: true,
		Details:    false,
		Timestamps: false,
	}

	logs, err := de.dockerClient.ContainerLogs(ctx, step.Metadata.UID, opts)
	if err != nil {
		return nil, err
	}
	return logs, err
	//rc, wc := io.Pipe()
	//
	//go func() {
	//	stdcopy.StdCopy(wc, wc, logs)
	//	logs.Close()
	//	wc.Close()
	//	rc.Close()
	//}()
	//return rc, nil
}
