package main

import (
	"github.com/workhorse/scheduler/pkg/scheduler"
)


func main() {

	sch := scheduler.NewScheduler()
	go sch.Run()

	for {

	}

	//var wg sync.WaitGroup
	//wg.Add(1)
	//
	//b := api.Builds{}
	//
	//go func() {
	//	b.Watch("http://localhost:8084/events", func(obj interface{}) {
	//		build := obj.(*coreapi.Build)
	//		nodebinding.BindBuildToNode(build)
	//		// log.Println("Build::::", build)
	//	})
	//}()
	//
	//go func() {
	//	b.WatchSteps("http://localhost:8084/events", func(obj interface{}) {
	//		buildStep := obj.(*coreapi.BuildStep)
	//		log.Println("Build-Step:::", buildStep)
	//		nodebinding.BindBuildStepToNode(buildStep)
	//	})
	//}()
	//
	//wg.Wait()
}
