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

func (ns *NodeInfoServer) Handle(response http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ns.list(response, r)
	} else if r.Method == http.MethodPost {
		ns.updateInfo(response, r)
	}
}

func (server *NodeInfoServer) updateInfo(response http.ResponseWriter, request *http.Request) {

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

func (server *NodeInfoServer) list(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		response.Header().Set("Content-Type", "application/json")
		nodeList, err := server.nis.ListNodes()
		if err != nil {
			log.Println(err)
			return
		}

		data, err := json.Marshal(nodeList)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = response.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
