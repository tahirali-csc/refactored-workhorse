package main

import (
	coreapi "github.com/workhorse/api"
	"github.com/workhorse/client/api"
	"log"

	//coreapi "github.com/workhorse/api"
	//"github.com/workhorse/client/api"
	"github.com/workhorse/worker/pkg/executor"
	"sync"

	//"log"
	//"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	ex := executor.NewExecutor()

	go func() {
		b := api.Builds{}

		b.WatchBuildStepNodeBinding("http://localhost:8084/events", func(obj interface{}) {
			buildStep := obj.(*coreapi.BuildStepNodeBinding)
			log.Println("Step::::", buildStep)

			ex.DataChannel <- buildStep
		})
	}()

	//executor.ExecuteStep(1)



	go func() {

		b := api.Builds{}

		b.WatchSteps("http://localhost:8084/events", func(obj interface{}) {
			//buildStep := obj.(*coreapi.BuildStep)
			//log.Println(buildStep)
		})
	}()



	ex.Run()

	wg.Wait()

}
