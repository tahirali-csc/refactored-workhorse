package main

import (
	"bufio"
	"context"
	"fmt"
	api2 "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	engine2 "github.com/workhorse/worker/pkg/engine"
	"github.com/workhorse/worker/pkg/engine/docker"
	"github.com/workhorse/worker/pkg/executor"
	"github.com/workhorse/worker/pkg/logstorage"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func executeStep(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		val, ok := request.URL.Query()["stepId"]
		if ok {
			stepId, _ := strconv.Atoi(val[0])
			b := api.Builds{}
			buildStep, _ := b.GetStep(stepId)

			file := "#!/bin/sh\n"
			for _, c := range buildStep.Commands {
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

			//response := runDockerContainer(jobDir, fileT, &step)
			//rd := bufio.NewReader(response)
			//
			//for {
			//	line, err := rd.ReadBytes('\n')
			//	if err != nil {
			//		break
			//	}
			//
			//	log.Print(string(line))
			//}

		}

	}
}

func RunStep() {
	jobDir, err := ioutil.TempDir("", "app")
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(jobDir)
	fileT, error := ioutil.TempFile(jobDir, "*.sh")
	if error != nil {
		log.Println(error)
	}

	err = fileT.Chmod(0777)
	if error != nil {
		log.Println(error)
	}

	contents := `#!/bin/sh
ls -la
date
sleep 10s
ls -la | grep 'me'`

	_, error = fileT.WriteString(contents)
	if error != nil {
		log.Println(error)
	}
	fileT.Close()

	log.Println(fileT.Name())

	step := engine2.Step{
		Metadata: engine2.Metadata{
			UID: fmt.Sprintf("%d", time.Now().Unix()),
		},
		Docker: &engine2.DockerStep{
			Args:    nil,
			Command: []string{"./" + fileT.Name()},
			Image:   "alpine:latest",
		},
		Volume: []*engine2.VolumeMount{
			{
				Source: jobDir,
				Target: jobDir,
			},
		},
	}
	engine, _ := docker.NewDockerEngine()
	err = engine.Create(context.Background(), nil, &step)
	if err != nil {
		log.Println(err)
	}

	err = engine.Start(context.Background(), nil, &step)
	if err != nil {
		log.Println(err)
	}

	reader, err := engine.Tail(context.Background(), nil, &step)
	linerReader := bufio.NewReader(reader)
	for {
		line, err := linerReader.ReadBytes('\n')
		if err != nil {
			break
		}
		log.Print(string(line))
	}
}

func runStep() {

	tempStepLogDir, err := ioutil.TempDir("", "setp")
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(tempStepLogDir)
	tempStepLogFile, error := ioutil.TempFile(tempStepLogDir, "*.log")
	if error != nil {
		log.Println(error)
	}

	log.Println("log location:", tempStepLogFile.Name())

	err = tempStepLogFile.Chmod(0777)
	if error != nil {
		log.Println(error)
	}

	dockerEngine, _ := docker.NewDockerEngine()
	fileLogStorage := logstorage.NewFileLogStore(tempStepLogFile)
	stepRunner := executor.NewStepRunner(dockerEngine, fileLogStorage)

	//Fake Build Step
	buildStep := &api2.BuildStep{}
	buildStep.Name = "build1"
	buildStep.Image = "alpine:latest"
	buildStep.Commands = []api2.BuildStepCommand{
		{
			Command: "ls -al",
			Id:      1,
			StepId:  1,
		},
		{
			Command: "sleep 10s && date",
			Id:      2,
			StepId:  1,
		},
	}
	//

	stepRunner.Run(buildStep)
}

func main() {

	runStep()

	//RunStep()

	//TODO: would review if there is an event arrived while waiting for the response from API server
	//nodeUpdater := nodeupdater.NewNodeUpdater()
	//nodeInfo, err := nodeUpdater.Register()
	//if err != nil {
	//	panic(err)
	//}
	//
	//go func() {
	//	be := executor.NewBuildExecutor()
	//	b := api.Builds{}
	//	for {
	//		b.WatchBuildNodeBinding("http://localhost:8084/events", func(obj interface{}) {
	//			buildNodeBinding := obj.(*coreapi.BuildNodeBinding)
	//
	//			if nodeInfo.Id == buildNodeBinding.NodeId {
	//				be.Execute(buildNodeBinding.BuildId)
	//			}
	//		})
	//	}
	//}()
	//
	//http.HandleFunc("/execStep", executeStep)
	//
	//http.ListenAndServe("localhost:8086", nil)

	//name, _ := os.Hostname()
	//log.Println(name)
	//
	//var wg sync.WaitGroup
	//wg.Add(1)
	//
	//ex := executor.NewExecutor()
	//
	//go func() {
	//	b := api.Builds{}
	//
	//	b.WatchBuildStepNodeBinding("http://localhost:8084/events", func(obj interface{}) {
	//		buildStep := obj.(*coreapi.BuildStepNodeBinding)
	//		log.Println("Step::::", buildStep)
	//
	//		ex.DataChannel <- buildStep
	//	})
	//}()
	//
	////executor.ExecuteStep(1)
	//
	//
	//
	//go func() {
	//
	//	b := api.Builds{}
	//
	//	b.WatchSteps("http://localhost:8084/events", func(obj interface{}) {
	//		//buildStep := obj.(*coreapi.BuildStep)
	//		//log.Println(buildStep)
	//	})
	//}()
	//
	//
	//
	//ex.Run()
	//
	//wg.Wait()

}
