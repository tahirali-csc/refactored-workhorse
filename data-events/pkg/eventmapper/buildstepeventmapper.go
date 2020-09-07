package eventmapper

import (
	"encoding/json"
	"github.com/workhorse/api"
	"github.com/workhorse/commons"
	"log"
)

type BuildStepEventObjectMapper struct {
}

func (mapper *BuildStepEventObjectMapper) Map(m map[string]interface{}) []byte {

	build := api.BuildStep{
		Id:      int(m["id"].(float64)),
		BuildId: int(m["build_id"].(float64)),
		Name:    m["name"].(string),
		Status:  m["status"].(string),
	}

	if m["created_ts"] != nil {
		build.CreatedTs = commons.ParseDBTime(m["created_ts"])
	}
	if m["end_ts"] != nil {
		build.EndTs = commons.ParseDBTime(m["end_ts"])
	}
	if m["start_ts"] != nil {
		build.StartTs = commons.ParseDBTime(m["start_ts"])
	}

	output, err := json.Marshal(build)
	if err != nil {
		log.Print(err)
	}
	return output
}
