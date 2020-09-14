package main

import "github.com/workhorse/worker/pkg/nodeupdater"

func main() {

	nodeUpdater := nodeupdater.NewNodeUpdater()
	nodeUpdater.Register()

	for {

	}

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
