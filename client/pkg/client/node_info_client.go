package client

import (
	"bytes"
	"encoding/json"
	"github.com/workhorse/api"
	"net/http"
)

type NodeInfoClient struct {
	apiClient *ApiClient
}

func NewNodeInfoClient(apClient *ApiClient) *NodeInfoClient {
	return &NodeInfoClient{
		apiClient: apClient,
	}
}

func (nodeClient *NodeInfoClient) Update(info *api.NodeInfo) error {

	obj, err := json.Marshal(info)
	if err != nil {
		return err
	}

	client := http.Client{}
	_, err = client.Post(nodeClient.apiClient.url+"/api/nodeinfo", "application/json", bytes.NewReader(obj))
	if err != nil {
		return err
	}

	return nil
}
