package client

import (
	"bytes"
	"encoding/json"
	"github.com/workhorse/api"
	"io/ioutil"
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

func (nodeClient *NodeInfoClient) List() ([]api.NodeInfo, error) {

	client := http.Client{}
	res, err := client.Get(nodeClient.apiClient.url + "/api/nodeinfo")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var nodeList []api.NodeInfo
	err = json.Unmarshal(data, &nodeList)
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}
