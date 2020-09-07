package eventmapper

import (
	"encoding/json"
	"github.com/workhorse/api"
	"log"
)

type BuildStepNodeBindingEventObjectMapper struct {
}

func (mapper *BuildStepNodeBindingEventObjectMapper) Map(m map[string]interface{}) []byte {

	build := api.BuildStepNodeBinding{
		Id:        int(m["id"].(float64)),
		StepId:    int(m["step_id"].(float64)),
		IpAddress: m["ip_address"].(string),
	}

	output, err := json.Marshal(build)
	if err != nil {
		log.Print(err)
	}
	return output
}
