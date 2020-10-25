package executor

import (
	api2 "github.com/workhorse/api"
	coreapi "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
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
			go orchestrator.RunBuild(binding.BuildId)
		}
	}
}

type step struct {
	Name     string   `yaml:"name"`
	Image    string   `yaml:"image"`
	Commands []string `yaml:"commands"`
}

type buildFile struct {
	Steps []step
}

func (orchestrator *BuildOrchestrator) RunBuild(buildId int) {
	orchestrator.buildStepMutex.Lock()
	orchestrator.buildStepEvents[buildId] = make(chan *coreapi.BuildStep)
	orchestrator.buildStepMutex.Unlock()

	b := api.Builds{}
	build, _ := b.GetBuild(buildId)

	tempDir, err := ioutil.TempDir("", "app")
	if err != nil {
		log.Fatal(err)
	}

	err = GitClone(tempDir, build.Project)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := ioutil.ReadFile(path.Join(tempDir, "build.yaml"))
	if err != nil {
		log.Println(err)
		return
	}

	bf := &buildFile{}
	err = yaml.Unmarshal(data, &bf)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(bf)

	var steps []api2.BuildStep
	for _, s := range bf.Steps {
		buildStep := api2.BuildStep{
			Name:    s.Name,
			Image:   s.Image,
			BuildId: buildId,
		}

		for _, c := range s.Commands {
			buildStep.Commands = append(buildStep.Commands, api2.BuildStepCommand{
				Command: c,
			})
		}

		steps = append(steps, buildStep)
	}

	//TODO: Build DAG
	builds := api.Builds{}
	builds.CreateBuildSteps(steps)

	build, _ = b.GetBuild(buildId)
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
