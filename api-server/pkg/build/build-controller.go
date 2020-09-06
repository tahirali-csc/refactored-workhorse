package build

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/workhorse/api"
)

type BuildController struct {
}

func (bc *BuildController) CreateBuild(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		mp := &buildInput{}
		json.Unmarshal(body, mp)

		build := api.Build{
			ProjectId: mp.ProjectID,
			Steps:     []api.BuildStep{},
		}

		for _, v := range mp.Steps {
			st := api.BuildStep{
				Image:    v.Image,
				Name:     v.Name,
				Commands: []api.BuildStepCommand{},
			}
			for _, c := range v.Commands {
				st.Commands = append(st.Commands, api.BuildStepCommand{
					Command: c,
				})
			}
			build.Steps = append(build.Steps, st)
		}

		buildService.StartNewBuild(build)
	}
}

func (bc *BuildController) BindToNode(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		mp := &api.BuildNodeBinding{}
		err = json.Unmarshal(body, mp)
		if err != nil {
			panic(err)
		}

		buildService.BindToNode(mp)
	}
}

type buildInput struct {
	ProjectID int64 `json:"projectId"`
	Steps     []struct {
		Image    string   `json:"image"`
		Name     string   `json:"name"`
		Commands []string `json:"commands"`
	} `json:"steps"`
}
