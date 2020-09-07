package nodebinding

import (
	"log"

	coreapi "github.com/workhorse/api"
	client "github.com/workhorse/client/api"
)

func BindBuildToNode(build *coreapi.Build) {

	selectedHost := "localhost"

	nodeBinding := coreapi.BuildNodeBinding{
		IpAddress: selectedHost,
		BuildId:   build.Id,
	}

	builds := client.Builds{}
	log.Println(builds)
	builds.BindToNode(nodeBinding)

}

func BindBuildStepToNode(step *coreapi.BuildStep) {

	selectedHost := "localhost"

	nodeBinding := coreapi.BuildStepNodeBinding{
		IpAddress: selectedHost,
		StepId:   step.Id,
	}

	builds := client.Builds{}
	log.Println(builds)
	builds.BindBuildStepToNode(nodeBinding)

}
