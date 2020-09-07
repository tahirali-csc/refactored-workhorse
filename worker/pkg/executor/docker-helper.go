package executor

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/workhorse/api"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ExecuteStep(stepId int) {
	step := api.BuildStep{
		Id:      stepId,
		BuildId: 1,
		Name:    "step1",
		Image:   "alpine:latest",
		Status:  "Pending",
		Commands: []api.BuildStepCommand{
			{
				Command: "ls -la",
			},
			{
				Command: "date",
			},
		},
	}

	file := "#!/bin/sh\n"
	for _, c := range step.Commands {
		file += c.Command + "\n"
	}

	jobDir, err := ioutil.TempDir("", "app")
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(jobDir)
	fileT, error := ioutil.TempFile(jobDir, "*.sh")
	if error != nil {
		log.Println(error)
	}

	log.Println(fileT.Name())
	err = fileT.Chmod(0777)
	if error != nil {
		log.Println(error)
	}

	_, error = fileT.WriteString(file)
	if error != nil {
		log.Println(error)
	}

	response := runDockerContainer(jobDir, fileT, &step)
	rd := bufio.NewReader(response)

	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			break
		}

		log.Print(string(line))
	}

}

func runDockerContainer(mountDir string, stepFile *os.File, step *api.BuildStep) io.ReadCloser{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	log.Printf("Pulling an image %s", step.Image)
	reader, err := cli.ImagePull(ctx, step.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: step.Image,
		Cmd:          []string{"/bin/sh", "-c", "." + stepFile.Name()},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		Entrypoint:   []string{},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: mountDir,
				Target: mountDir,
			},
		},
	}, nil, "")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}

	return out


}
