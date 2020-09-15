package nodeupdater

import (
	"github.com/workhorse/api"
	"github.com/workhorse/client/pkg/client"
	"github.com/workhorse/commons"
	"time"
)

type NodeUpdater struct {
	nodeInfoClient *client.NodeInfoClient
}

func NewNodeUpdater() *NodeUpdater {
	apiClient := client.ApiClient{}
	apiClient.Init("http://localhost:8081")

	return &NodeUpdater{
		nodeInfoClient: apiClient.GetNodeInfoClient(),
	}
}

func (updater *NodeUpdater) Register() {
	updater.update()
	go updater.heartBeat()
}

func (updater *NodeUpdater) heartBeat() {
	timer := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-timer.C:
			updater.update()
		}
	}
}

func (updater *NodeUpdater) update() error {
	hostName, err := commons.GetHostname()
	if err != nil {
		return err
	}

	nodeInfo := &api.NodeInfo{
		Name:            hostName,
		LastHeartBeatTS: time.Now(),
	}

	return updater.nodeInfoClient.Update(nodeInfo)
}
