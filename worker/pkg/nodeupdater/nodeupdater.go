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

func (updater *NodeUpdater) Register() (*api.NodeInfo, error) {
	go updater.heartBeat()
	return updater.update()
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

func (updater *NodeUpdater) update() (*api.NodeInfo, error) {
	hostName, err := commons.GetHostname()
	if err != nil {
		return nil, err
	}

	nodeInfo := &api.NodeInfo{
		Name:            hostName,
		LastHeartBeatTS: time.Now(),
	}

	return updater.nodeInfoClient.Update(nodeInfo)
}
