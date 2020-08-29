package eventmapper

import (
	"encoding/json"
	"net/http"

	"gihtub.com/workhorse/data-events/pkg/listener"

	"github.com/r3labs/sse"
)

type eventInfo struct {
	Table  string                 `json:"table"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type HTTPHandler func(http.ResponseWriter, *http.Request)

type DBEventToSSEStreamMapper struct {
	tableStreamMap   map[string]eventMappingInfo
	dbEventsListener *listener.DatabaseEventsListener
	sseServer        *sse.Server
	HTTPHandler      HTTPHandler
}

func New(dbEventsListener *listener.DatabaseEventsListener) *DBEventToSSEStreamMapper {
	mapper := &DBEventToSSEStreamMapper{
		tableStreamMap:   make(map[string]eventMappingInfo),
		dbEventsListener: dbEventsListener,
		sseServer:        sse.New(),
	}
	mapper.initEventToSSEMapping()
	mapper.HTTPHandler = mapper.sseServer.HTTPHandler
	return mapper
}

func (mapper *DBEventToSSEStreamMapper) initEventToSSEMapping() {
	mapper.tableStreamMap["build"] = eventMappingInfo{Stream: "build", Mapper: &BuildEventObjectMapper{}}
	// mapper.tableStreamMap["build_jobs"] = eventMappingInfo{Stream: "build-jobs", Mapper: &BuildEventObjectMapper{}}

	for _, value := range mapper.tableStreamMap {
		mapper.sseServer.CreateStream(value.Stream)
	}
}

func (mapper *DBEventToSSEStreamMapper) WatchEvents() {
	for {
		select {
		case event := <-mapper.dbEventsListener.EventChannel:
			ei := &eventInfo{}
			json.Unmarshal([]byte(event), ei)

			// record, err := json.Marshal(ei.Data)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			if eventStream, ok := mapper.tableStreamMap[ei.Table]; ok {
				data := eventStream.Mapper.Map(ei.Data)

				mapper.sseServer.Publish(eventStream.Stream, &sse.Event{
					Data: data,
				})
			}
		}
	}
}

type eventMappingInfo struct {
	Stream string
	Mapper EventObjectMapper
}
