package project

import (
	"encoding/json"
	"github.com/workhorse/api"
	"io/ioutil"
	"log"
	"net/http"
)

type ProjectServer struct {
	service ProjectService
}

func NewProjectServer() *ProjectServer {
	return &ProjectServer{}
}

func (ps *ProjectServer) Handle(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		res, err := ps.createProject(request)
		if err != nil {
			log.Println(err)
			http.Error(response, "Unable to save project", 500)
			return
		}

		data, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			http.Error(response, "Unable to convert an input object", 500)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	}
}

func (ps *ProjectServer) createProject(request *http.Request) (*api.Project, error) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	project := &api.Project{}
	err = json.Unmarshal(body, project)
	if err != nil {
		return nil, err
	}

	res, err := ps.service.Create(project)
	return res, err

}
