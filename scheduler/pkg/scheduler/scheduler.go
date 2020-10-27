package scheduler

import (
	"github.com/workhorse/api"
	client2 "github.com/workflow/client/api"
	"github.com/workhorse/client/pkg/client"
	"github.com/workhorse/scheduler/pkg/nodebinding"
	"github.com/workhorse/scheduler/pkg/schedulingalgos/buildscheduling"
	"log"
	"sync"
	"time"
)

type Scheduler struct {
	nodeList     []api.NodeInfo
	nodeListLock *sync.RWMutex
	bq           *BuildQueue
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		nodeListLock: &sync.RWMutex{},
		bq:           NewBuildQueue(),
	}
}

func (sch *Scheduler) Run() {
	go sch.listNodes()

	go func() {
		b := client2.Builds{}
		for {
			b.Watch("http://localhost:8084/events", func(obj interface{}) {
				log.Println("Got event :::", obj)
				build := obj.(*api.Build)
				sch.bq.Add(build)
			})
		}
	}()

	d, _ := time.ParseDuration("10s")
	timer := time.NewTicker(d)
	parallelBuilds := 3
	schStrategy := buildscheduling.NewRoundRobinScheduling(&sch.nodeList)

	for {
		select {
		case <-timer.C:
			for i := 0; i < parallelBuilds; i++ {
				if sch.bq.Length() > 0 && len(sch.nodeList) > 0 {
					//log.Println("Node Count::", len(sch.nodeList))
					build := sch.bq.Pop()
					node := schStrategy.Pick(build)
					nodebinding.BindBuildToNode(build, node)
				}
			}
		}
	}
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

	var newNodes []api.NodeInfo
	for _, n := range nodeList {
		for i := range sch.nodeList {
			if sch.nodeList[i].Id == n.Id {
				sch.nodeList[i].Name = n.Name
				sch.nodeList[i].LastHeartBeatTS = n.LastHeartBeatTS
			} else {
				newNodes = append(newNodes, n)
			}
		}
	}

	sch.nodeList = append(sch.nodeList, newNodes...)
	//log.Println("Nodes:::", sch.nodeList)
}
