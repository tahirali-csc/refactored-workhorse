package buildscheduling

import (
	"github.com/workhorse/api"
)

type Scheduling interface {
	Pick(build *api.Build) api.NodeInfo
}

type RoundRobinScheduling struct {
	index int
	nodes *[]api.NodeInfo
}

func (rr *RoundRobinScheduling) Pick(_ *api.Build) api.NodeInfo {
	nodes := *rr.nodes
	//log.Println("RoundRobinScheduling Node Count::", len(nodes))

	node := nodes[rr.index%len(nodes)]
	rr.index++
	return node
}

func NewRoundRobinScheduling(nodes *[]api.NodeInfo) Scheduling {
	return &RoundRobinScheduling{nodes: nodes}
}