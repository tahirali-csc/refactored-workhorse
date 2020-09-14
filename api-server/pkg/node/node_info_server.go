package node

import (
	"encoding/json"
	"github.com/workhorse/api"
	"io/ioutil"
	"log"
	"net/http"
)

type NodeInfoServer struct {
	nis NodeInfoService
}

func (server *NodeInfoServer) UpdateInfo(response http.ResponseWriter, request *http.Request)  {

	if request.Method == http.MethodPost {
		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}

		info := &api.NodeInfo{}
		err = json.Unmarshal(body, info)
		if err != nil {
			log.Println(err)
			return
		}

		err = server.nis.UpdateNode(info)
		if err != nil {
			log.Println(err)
			return
		}
	}
}