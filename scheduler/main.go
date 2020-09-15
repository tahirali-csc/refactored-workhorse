package main

import (
	"github.com/workhorse/api"
	client2 "github.com/workhorse/client/api"
	"github.com/workhorse/client/pkg/client"
	"log"
	"sync"
	"time"
)

type Scheduler struct {
	nodeList     []api.NodeInfo
	nodeListLock *sync.RWMutex
}

func (sch *Scheduler) Run() {
	go sch.listNodes()

	b := client2.Builds{}
	b.Watch("http://localhost:8084/events", func(obj interface{}) {
		build := obj.(*api.Build)
		log.Println(build)
		//nodebinding.BindBuildToNode(build)
	})

}

func (sch *Scheduler) listNodes() {
	apiClient := client.ApiClient{}
	apiClient.Init("http://localhost:8081/")
	nodeInfoClient := apiClient.GetNodeInfoClient()

	timer := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-timer.C:
			nodeList, err := nodeInfoClient.List()
			if err != nil {
				log.Println(err)
			} else {
				if len(sch.nodeList) == 0 {
					sch.nodeList = nodeList
				} else {
					//log.Println("Got::", nodeList)
					sch.syncNodeList(nodeList)
				}
			}
		}
	}
}

func (sch *Scheduler) syncNodeList(nodeList []api.NodeInfo) {
	defer sch.nodeListLock.Unlock()
	sch.nodeListLock.Lock()

	for _, n := range nodeList {
		for i := range sch.nodeList {
			if sch.nodeList[i].Id == n.Id {
				sch.nodeList[i].Name = n.Name
				sch.nodeList[i].LastHeartBeatTS = n.LastHeartBeatTS
				//log.Println(sch.nodeList[i])
			}
		}
	}
	//log.Println(sch.nodeList)
	//log.Println()
}

func main() {

	sch := Scheduler{
		nodeListLock: &sync.RWMutex{},
	}

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
