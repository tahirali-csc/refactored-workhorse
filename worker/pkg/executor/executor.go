package executor

import (
	"github.com/workhorse/api"
	client "github.com/workhorse/client/api"
	"log"
)

type Executor struct {
	DataChannel chan interface{}
}

func NewExecutor() *Executor {
	return &Executor{DataChannel: make(chan interface{})}
}

func (ex *Executor) Run() {

	build := api.Build{
		Id:        38,
		Status:    "Pending",
		ProjectId: 1,
		Steps: []api.BuildStep{
			{
				Id:     29,
				Image:  "alpine:latest",
				Status: "Pending",
				Commands: []api.BuildStepCommand{
					{
						Command: "ls -la",
					},
					{
						Command: "date",
					},
				},
			},
			{
				Id:     30,
				Image:  "alpine:latest",
				Status: "Pending",
				Commands: []api.BuildStepCommand{
					{
						Command: "echo 'Hello World'",
					},
					{
						Command: "echo 'Hello J'",
					},
				},
			},
		},
	}

	builds := client.Builds{}
	log.Println(len(build.Steps))
	for _, step := range build.Steps {
		log.Println("Executing step......")
		builds.UpdateBuildStepStatus(step.Id, "TrySchedule")

	loop:
		for {
			select {
			case msg := <-ex.DataChannel:
				bindingn := msg.(*api.BuildStepNodeBinding)
				if bindingn.StepId == step.Id {
					log.Println("Yeah!!!", bindingn)
					ExecuteStep(step.Id)
					break loop
				}
			}
		}





		log.Println("Goig ")
	}

}
