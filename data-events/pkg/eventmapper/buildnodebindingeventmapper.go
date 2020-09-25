package eventmapper

import (
	"encoding/json"
	"github.com/workhorse/api"
	"log"
)

type BuildNodeBindingEventMapper struct {
}

func (mapper *BuildNodeBindingEventMapper) Map(m map[string]interface{}) []byte {

	build := api.BuildNodeBinding{
		Id:      int(m["id"].(float64)),
		BuildId: int(m["build_id"].(float64)),
		NodeId:  int(m["node_id"].(float64)),
	}

	output, err := json.Marshal(build)
	if err != nil {
		log.Print(err)
	}
	return output
}
