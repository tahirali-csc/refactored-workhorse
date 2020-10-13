package buildstep

import (
	api2 "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	"github.com/workhorse/worker/pkg/engine/docker"
	"github.com/workhorse/worker/pkg/executor"
	"github.com/workhorse/worker/pkg/logstorage"
	"io/ioutil"
	"log"
)

type StepManager struct {
}

func NewStepManager() *StepManager {
	return &StepManager{}
}

func (manager *StepManager) Run(stepId int) {

	b := api.Builds{}
	buildStep, _ := b.GetStep(stepId)

	tempStepLogDir, err := ioutil.TempDir("", "setp")
	if err != nil {
		log.Fatal(err)
	}

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

	buildStep.Status = "Starting"
	//b.UpdateBuildStep(stepId, "Starting")
	b.UpdateBuildStep(buildStep)


	err = stepRunner.Run(buildStep)
	if err != nil {
		log.Println(err)
	}

	//b.UpdateBuildStepStatus(stepId, "Finished")
	buildStep.Status = "Finished"
	//b.UpdateBuildStep(stepId, "Starting")
	buildStep.LogInfo = make(api2.LogStorageProperties)
	buildStep.LogInfo["type"] = "file"
	buildStep.LogInfo["path"] = tempStepLogFile.Name()
	b.UpdateBuildStep(buildStep)

}
