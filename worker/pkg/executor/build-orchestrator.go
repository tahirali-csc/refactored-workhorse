package executor

import (
	coreapi "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	"log"
	"sync"
)

type BuildOrchestrator struct {
	buildBindingChan chan *coreapi.BuildNodeBinding
	buildStepEvents  map[int]chan *coreapi.BuildStep
	buildStepMutex   sync.RWMutex
}

func NewBuildOrchestrator() *BuildOrchestrator {
	return &BuildOrchestrator{
		buildBindingChan: make(chan *coreapi.BuildNodeBinding),
		buildStepEvents:  make(map[int]chan *coreapi.BuildStep),
		buildStepMutex:   sync.RWMutex{},
	}
}

func (orchestrator *BuildOrchestrator) Start() {
	go orchestrator.watchBuildNodeBinding()
	go orchestrator.watchBuildStepStatus()

	for {
		select {
		case binding := <-orchestrator.buildBindingChan:
			go orchestrator.runBuild(binding.BuildId)
		}
	}
}

func (orchestrator *BuildOrchestrator) runBuild(buildId int) {
	orchestrator.buildStepMutex.Lock()
	orchestrator.buildStepEvents[buildId] = make(chan *coreapi.BuildStep)
	orchestrator.buildStepMutex.Unlock()

	b := api.Builds{}
	build, _ := b.GetBuild(buildId)

	for _, s := range build.Steps {
		log.Println("Starting build step", s.Id)
		err := b.RunStep(s.Id)
		if err != nil {
			log.Println(err)
		}

	loop:
		for {
			select {
			case bs := <-orchestrator.buildStepEvents[buildId]:
				if bs.BuildId == buildId && bs.Status == "Finished" {
					log.Println("Got build step status:::", bs.Status)
					break loop
				}
			}
		}
	}

}

func (orchestrator *BuildOrchestrator) watchBuildNodeBinding() {
	b := api.Builds{}
	b.WatchBuildNodeBinding("http://localhost:8084/events", func(obj interface{}) {
		orchestrator.buildBindingChan <- obj.(*coreapi.BuildNodeBinding)
	})
}

func (orchestrator *BuildOrchestrator) watchBuildStepStatus() {
	b := api.Builds{}
	b.WatchSteps("http://localhost:8084/events", func(obj interface{}) {
		buildStep := obj.(*coreapi.BuildStep)
		orchestrator.buildStepMutex.RLock()
		orchestrator.buildStepEvents[buildStep.BuildId] <- buildStep
		orchestrator.buildStepMutex.RUnlock()
	})
}
