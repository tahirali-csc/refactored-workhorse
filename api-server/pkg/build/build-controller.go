package build

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func (bc *BuildController) UpdateBuildStepStatus(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		bs := &api.BuildStep{}

		//mp := &buildStepStatus{}
		//err = json.Unmarshal(body, mp)
		//if err != nil {
		//	panic(err)
		//}
		//
		//buildService.UpdateBuildStepStatus(mp.StepId, mp.Status)

		err = json.Unmarshal(body, bs)
		if err != nil {
			panic(err)
		}
		buildService.UpdateBuildStep(bs)

	}
}

func (bc *BuildController) BindingBuildStepToNode(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		mp := &api.BuildStepNodeBinding{}
		err = json.Unmarshal(body, mp)
		if err != nil {
			panic(err)
		}

		buildService.BindBuildStepToNode(mp)
	}
}

func (bc *BuildController) GetBuild(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		buildService := BuildService{}

		val, ok := request.URL.Query()["buildId"]
		if ok {
			buildId, _ := strconv.Atoi(val[0])
			steps, _ := buildService.GetBuild(buildId)
			dat, err := json.Marshal(steps)
			if err != nil {
				log.Println(err)
				return
			}
			response.Write(dat)
		}
	}
}

func (bc *BuildController) GetStep(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		buildService := BuildService{}

		val, ok := request.URL.Query()["stepId"]
		if ok {
			stepId, _ := strconv.Atoi(val[0])
			steps, _ := buildService.GetStep(stepId)
			dat, err := json.Marshal(steps)
			if err != nil {
				log.Println(err)
				return
			}
			response.Write(dat)
		}
	}
}

func (bc *BuildController) Patch(response http.ResponseWriter, request *http.Request){
	if request.Method == "PATCH"{
		bodyData , _ := ioutil.ReadAll(request.Body)

		patchData := make(map[string]interface{})
		err := json.Unmarshal(bodyData, &patchData)
		if err != nil {
			log.Println(err)
			return
		}



		id, _ := patchData["id"].(float64)
		delete(patchData, "id")

		buildService := BuildService{}
		buildService.PatchBuildStep(int(id), patchData)
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

type buildStepStatus struct {
	StepId int
	Status string
}
