package main

import (
	"sync"

	"github.com/workhorse/scheduler/pkg/nodebinding"

	coreapi "github.com/workhorse/api"
	"github.com/workhorse/client/api"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	b := api.Builds{}

	go func() {
		b.Watch("http://localhost:8084/events", func(obj interface{}) {
			build := obj.(*coreapi.Build)
			nodebinding.BindBuildToNode(build)
			// log.Println("Build::::", build)
		})
	}()

	wg.Wait()
}
