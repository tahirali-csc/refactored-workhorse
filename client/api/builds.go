package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/r3labs/sse"
	"github.com/workhorse/api"
)

type WatchHandler func(obj interface{})

type Builds struct {
}

func (b *Builds) Watch(url string, handler WatchHandler) error {

	client := sse.NewClient(url)

	err := client.Subscribe("build", func(msg *sse.Event) {
		log.Println("Event ID:", client.EventID)
		build := &api.Build{}
		json.Unmarshal(msg.Data, build)
		handler(build)
	})

	return err
}

func (b *Builds) WatchSteps(url string, handler WatchHandler) error {

	client := sse.NewClient(url)

	err := client.Subscribe("build-steps", func(msg *sse.Event) {
		build := &api.BuildStep{}
		json.Unmarshal(msg.Data, build)
		handler(build)
	})

	return err
}

func (b *Builds) WatchBuildStepNodeBinding(url string, handler WatchHandler) error {

	client := sse.NewClient(url)

	err := client.Subscribe("build-steps-node-binding", func(msg *sse.Event) {
		build := &api.BuildStepNodeBinding{}
		//log.Println(msg)
		json.Unmarshal(msg.Data, build)
		handler(build)
	})

	return err
}

func (b *Builds) WatchBuildNodeBinding(url string, handler WatchHandler) error {

	client := sse.NewClient(url)

	err := client.Subscribe("build-node-binding", func(msg *sse.Event) {
		build := &api.BuildNodeBinding{}
		//log.Println(msg)
		json.Unmarshal(msg.Data, build)
		handler(build)
	})

	return err
}

func (b *Builds) BindToNode(binding api.BuildNodeBinding) {
	client := http.Client{}

	dt, _ := json.Marshal(binding)

	res, err := client.Post("http://localhost:8081/buildbinding", "application/json", bytes.NewReader(dt))
	log.Println(res)
	if err != nil {
		log.Println("Error", err)
	}

}

//func (b *Builds) UpdateBuildStepStatus(stepId int, status string) {
//	client := http.Client{}
//
//	st := make(map[string]interface{})
//	st["stepId"] = stepId
//	st["status"] = status
//
//	dt, _ := json.Marshal(st)
//
//	res, err := client.Post("http://localhost:8081/updatestepstatus", "application/json", bytes.NewReader(dt))
//	log.Println(res)
//	if err != nil {
//		log.Println("Error", err)
//	}
//
//}
func (b *Builds) UpdateBuildStep(step *api.BuildStep) error {
	dt, _ := json.Marshal(step)

	client := http.Client{}
	_, err := client.Post("http://localhost:8081/updatestepstatus", "application/json", bytes.NewReader(dt))
	return err
}

func (b *Builds) BindBuildStepToNode(binding api.BuildStepNodeBinding) {
	client := http.Client{}

	dt, _ := json.Marshal(binding)

	res, err := client.Post("http://localhost:8081/updatebuildstepbinding", "application/json", bytes.NewReader(dt))
	log.Println(res)
	if err != nil {
		log.Println("Error", err)
	}

}

func (b *Builds) GetBuild(buildId int) (*api.Build, error) {
	client := http.Client{}

	res, err := client.Get(fmt.Sprintf("http://localhost:8081/api/build?buildId=%d", buildId))
	//log.Println(res)
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	build := &api.Build{}
	err = json.Unmarshal(dat, build)
	return build, err
}

func (s *Builds) GetStep(stepId int) (*api.BuildStep, error) {
	client := http.Client{}

	res, err := client.Get(fmt.Sprintf("http://localhost:8081/getstep?stepId=%d", stepId))
	//log.Println(res)
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	build := &api.BuildStep{}
	err = json.Unmarshal(dat, build)
	return build, err
}

func (b *Builds) RunStep(stepId int) error {
	client := http.Client{}

	_, err := client.Post(fmt.Sprintf("http://localhost:8086/runstep?stepId=%d", stepId), "", nil)
	//log.Println(res)
	if err != nil {
		return err
	}

	//dat, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return nil, err
	//}

	//build := &api.Build{}
	//err = json.Unmarshal(dat, build)
	return err
}

func (b *Builds) CreateBuildSteps(steps []api.BuildStep) error {
	dt, _ := json.Marshal(steps)
	client := http.Client{}

	_, err := client.Post("http://localhost:8081/api/buildsteps", "application/json", bytes.NewReader(dt))
	//log.Println(res)
	if err != nil {
		return err
	}

	//dat, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return nil, err
	//}

	//build := &api.Build{}
	//err = json.Unmarshal(dat, build)
	return err
}
