package build

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/workhorse/api"
)

type BuildController struct {
}

func (bc *BuildController) createBuild(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		mp := &buildInput{}
		json.Unmarshal(body, mp)

		build := api.Build{
			//ProjectId: mp.ProjectID,
			Project: api.Project{
				Id:         int(mp.ProjectID),
				Name:       "",
				PrivateKey: "",
				CloneURL:   "",
			},
			Steps: []api.BuildStep{},
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

func (bc *BuildController) getBuild(response http.ResponseWriter, request *http.Request) {
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

func (bc *BuildController) Patch(response http.ResponseWriter, request *http.Request) {
	if request.Method == "PATCH" {
		bodyData, _ := ioutil.ReadAll(request.Body)

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

func (bc *BuildController) TailLogStep(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		val, ok := request.URL.Query()["stepId"]
		if ok {
			buildService := BuildService{}
			stepId, _ := strconv.Atoi(val[0])
			stepInfo, _ := buildService.GetStep(stepId)

			url := fmt.Sprintf("http://%s:8086/stream/step?stepId=%d", stepInfo.Node.Name, stepInfo.Id)

			log.Println("Stream URL::", url)
			client := http.Client{}
			res, err := client.Get(url)
			if err != nil {
				log.Println(err)
				return
			}

			h := w.Header()
			h.Set("Content-Type", "text/event-stream")
			h.Set("Cache-Control", "no-cache")
			h.Set("Connection", "keep-alive")
			h.Set("X-Accel-Buffering", "no")

			f, ok := w.(http.Flusher)
			if !ok {
				return
			}

			io.WriteString(w, ": ping\n\n")
			f.Flush()

			enc := json.NewEncoder(w)
			reader := bufio.NewReader(res.Body)

			for {
				line, err := reader.ReadBytes('\n')
				if err != nil {
					f.Flush()
					break
				}

				io.WriteString(w, "data: ")
				enc.Encode(string(line))
				io.WriteString(w, "\n\n")
				f.Flush()
			}

		}
	}
}

func (bc *BuildController) Handle(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		bc.getBuild(writer, request)
	} else if request.Method == http.MethodPost {
		bc.createBuild(writer, request)
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

func (bc *BuildController) HandleBuildStep(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		buildService := BuildService{}

		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		var bs []api.BuildStep
		json.Unmarshal(body, &bs)

		//build := api.Build{
		//	//ProjectId: mp.ProjectID,
		//	Project: api.Project{
		//		Id:         int(mp.ProjectID),
		//		Name:       "",
		//		PrivateKey: "",
		//		CloneURL:   "",
		//	},
		//	Steps: []api.BuildStep{},
		//}
		//
		//for _, v := range mp.Steps {
		//	st := api.BuildStep{
		//		Image:    v.Image,
		//		Name:     v.Name,
		//		Commands: []api.BuildStepCommand{},
		//	}
		//	for _, c := range v.Commands {
		//		st.Commands = append(st.Commands, api.BuildStepCommand{
		//			Command: c,
		//		})
		//	}
		//	build.Steps = append(build.Steps, st)
		//}

		buildService.CreateBuildSteps( bs)
	}
}
