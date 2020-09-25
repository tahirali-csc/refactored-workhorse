package executor

import (
	"fmt"
	api2 "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	"log"
	"net/http"
)

type BuildStepExecutor struct {
}

func (bse *BuildStepExecutor) Run(step api2.BuildStep) {
	client := http.Client{}
	client.Post(fmt.Sprintf("http://localhost:8086/execStep?stepId=%d", step.Id),
		"application/json", nil)
}

type BuildExecutor struct {
}

func (be *BuildExecutor) Execute(buildId int) error {
	bs := api.Builds{}
	build, err := bs.GetBuild(buildId)
	if err != nil {
		log.Println(err)
		return err
	}

	bse := &BuildStepExecutor{}
	for _, step := range build.Steps {
		bse.Run(step)
	}

	log.Println(build)
	return nil

}

func NewBuildExecutor() *BuildExecutor {
	return &BuildExecutor{}
}

//type Executor struct {
//	DataChannel chan interface{}
//}
//
//func NewExecutor() *Executor {
//	return &Executor{DataChannel: make(chan interface{})}
//}
//
//func (ex *Executor) Run() {
//
//	build := api.Build{
//		Id:        38,
//		Status:    "Pending",
//		ProjectId: 1,
//		Steps: []api.BuildStep{
//			{
//				Id:     29,
//				Image:  "alpine:latest",
//				Status: "Pending",
//				Commands: []api.BuildStepCommand{
//					{
//						Command: "ls -la",
//					},
//					{
//						Command: "date",
//					},
//				},
//			},
//			{
//				Id:     30,
//				Image:  "alpine:latest",
//				Status: "Pending",
//				Commands: []api.BuildStepCommand{
//					{
//						Command: "echo 'Hello World'",
//					},
//					{
//						Command: "echo 'Hello J'",
//					},
//				},
//			},
//		},
//	}
//
//	builds := client.Builds{}
//	log.Println(len(build.Steps))
//	for _, step := range build.Steps {
//		log.Println("Executing step......")
//		builds.UpdateBuildStepStatus(step.Id, "TrySchedule")
//
//	loop:
//		for {
//			select {
//			case msg := <-ex.DataChannel:
//				bindingn := msg.(*api.BuildStepNodeBinding)
//				if bindingn.StepId == step.Id {
//					log.Println("Yeah!!!", bindingn)
//					ExecuteStep(step.Id)
//					break loop
//				}
//			}
//		}
//
//
//
//
//
//		log.Println("Goig ")
//	}
//
//}
