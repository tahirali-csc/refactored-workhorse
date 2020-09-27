package executor

import (
	"bufio"
	"context"
	"fmt"
	"github.com/workhorse/api"
	engine2 "github.com/workhorse/worker/pkg/engine"
	"github.com/workhorse/worker/pkg/logstorage"
	"io/ioutil"
	"log"
	"time"
)

type StepRunner struct {
	engine   engine2.Engine
	logStore logstorage.LogStore
}

func NewStepRunner(engine engine2.Engine, logStore logstorage.LogStore) *StepRunner {
	return &StepRunner{
		engine:   engine,
		logStore: logStore,
	}
}

func (se *StepRunner) Run(buildStep *api.BuildStep) error {

	file := "#!/bin/sh\n"
	for _, c := range buildStep.Commands {
		file += c.Command + "\n"
	}

	tempBuildDir, err := ioutil.TempDir("", "build")
	if err != nil {
		return err
	}

	tempBuildFile, error := ioutil.TempFile(tempBuildDir, "*.sh")
	if error != nil {
		return err
	}

	err = tempBuildFile.Chmod(0777)
	if error != nil {
		return err
	}

	_, error = tempBuildFile.WriteString(file)
	if error != nil {
		return err
	}

	step := engine2.Step{
		Metadata: engine2.Metadata{
			UID: fmt.Sprintf("%d", time.Now().Unix()),
		},
		Docker: &engine2.DockerStep{
			Args:    nil,
			Command: []string{"./" + tempBuildFile.Name()},
			Image:   buildStep.Image,
		},
		Volume: []*engine2.VolumeMount{
			{
				Source: tempBuildDir,
				Target: tempBuildDir,
			},
		},
	}
	ctx := context.Background()

	err = se.engine.Create(ctx, nil, &step)
	if err != nil {
		log.Println(err)
	}

	err = se.engine.Start(ctx, nil, &step)
	if err != nil {
		log.Println(err)
	}

	reader, err := se.engine.Tail(ctx, nil, &step)

	linerReader := bufio.NewReader(reader)
	for {
		line, err := linerReader.ReadBytes('\n')
		if err != nil {
			break
		}
		se.logStore.Write(ctx, int64(buildStep.Id), line)
	}

	return nil
}
