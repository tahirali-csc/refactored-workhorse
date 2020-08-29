package eventmapper

import (
	"encoding/json"
	"log"

	"github.com/workhorse/api"
	"github.com/workhorse/commons"
)

type EventObjectMapper interface {
	Map(obj map[string]interface{}) []byte
}

type BuildEventObjectMapper struct {
}

func (mapper *BuildEventObjectMapper) Map(m map[string]interface{}) []byte {

	build := api.Build{
		Id:        int(m["id"].(float64)),
		ProjectId: int(m["project_id"].(float64)),
		Status:    string(m["status"].(string)),
		CreatedTs: commons.ParseDBTime(m["created_ts"]),
		EndTs:     commons.ParseDBTime(m["end_ts"]),
		StartTs:   commons.ParseDBTime(m["start_ts"]),
	}

	output, err := json.Marshal(build)
	if err != nil {
		log.Print(err)
	}
	return output
}
