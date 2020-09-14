package client

type ApiClient struct {
	url string
}

func (client *ApiClient) Init(url string) {
	client.url = url
}

func (client *ApiClient) GetNodeInfoClient() *NodeInfoClient {
	return NewNodeInfoClient(client)
}
