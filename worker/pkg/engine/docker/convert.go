package docker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/workhorse/worker/pkg/engine"
)

/*
Image: step.Image,
		Cmd:          []string{"/bin/sh", "-c", "." + stepFile.Name()},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		Entrypoint:   []string{},
*/
func toConfig(spec *engine.Spec, step *engine.Step) *container.Config {
	config := &container.Config{
		Image:        step.Docker.Image,
		Labels:       step.Metadata.Labels,
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		OpenStdin:    false,
		StdinOnce:    false,
		ArgsEscaped:  false,
	}

	if len(step.Docker.Args) != 0 {
		config.Cmd = step.Docker.Args
	}
	if len(step.Docker.Command) != 0 {
		config.Entrypoint = step.Docker.Command
	}

	return config
}

func toHostConfig(spec *engine.Spec, step *engine.Step) *container.HostConfig {
	var mounts []mount.Mount
	for _, m := range step.Volume {
		mounts = append(mounts, mount.Mount{
			Target: m.Target,
			Source: m.Source,
			Type:   mount.TypeBind,
		})
	}
	config := &container.HostConfig{
		Mounts: mounts,
	}
	return config
}
