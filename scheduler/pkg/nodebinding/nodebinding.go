package nodebinding

import (
	"github.com/workhorse/api"
	"log"

	coreapi "github.com/workhorse/api"
	client "github.com/workhorse/client/api"
)

func BindBuildToNode(build *coreapi.Build, node api.NodeInfo) {

	nodeBinding := coreapi.BuildNodeBinding{
		NodeId: node.Id,
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
