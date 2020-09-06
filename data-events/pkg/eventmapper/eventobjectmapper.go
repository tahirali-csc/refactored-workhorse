package eventmapper

import (
	"encoding/json"
	"github.com/workhorse/commons"
	"log"

	"github.com/workhorse/api"
)

type EventObjectMapper interface {
	Map(obj map[string]interface{}) []byte
}

type BuildEventObjectMapper struct {
}

func (mapper *BuildEventObjectMapper) Map(m map[string]interface{}) []byte {

	build := api.Build{
		Id:        int(m["id"].(float64)),
		ProjectId: int64(m["project_id"].(float64)),
		Status:    string(m["status"].(string)),
		//CreatedTs: commons.ParseDBTime(m["created_ts"]),
		//EndTs:     commons.ParseDBTime(m["end_ts"]),
		//StartTs:   commons.ParseDBTime(m["start_ts"]),
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
