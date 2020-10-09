package buildstep

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/workhorse/worker/pkg/engine"
	"github.com/workhorse/worker/pkg/engine/docker"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	//sseServer *sse.Server
}

func NewServer() *Server {
	//srv := sse.New()
	return &Server{
		//sseServer: srv,
	}
}

func (server *Server) HandleRunStep(reponse http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		val, ok := request.URL.Query()["stepId"]
		if ok {
			stepManager := NewStepManager()
			stepId, _ := strconv.Atoi(val[0])
			go stepManager.Run(stepId)
		}
	}
}

func (server *Server) HandleLogStream(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		val, ok := request.URL.Query()["stepId"]
		if ok {
			stepId, _ := strconv.Atoi(val[0])
			//streamId := fmt.Sprintf("Step: %d", stepId)

			step := &engine.Step{
				Metadata: engine.Metadata{
					UID: fmt.Sprintf("%d", stepId),
				},
			}

			dockerEngine, _ := docker.NewDockerEngine()
			reader, err := dockerEngine.Tail(context.Background(), nil, step)
			if err != nil {
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

			//if !server.sseServer.StreamExists(streamId) {
			//	server.sseServer.CreateStream(streamId)
			//}

			//go server.sseServer.HTTPHandler(reponse, request)

			enc := json.NewEncoder(w)
			logReader := bufio.NewReader(reader)
			for {
				line, err := logReader.ReadBytes('\n')
				if err != nil {
					f.Flush()
					break
				}
				io.WriteString(w, "data: ")
				enc.Encode(string(line))
				io.WriteString(w, "\n\n")
				f.Flush()

				//log.Print(string(line))

				//server.sseServer.Publish(streamId, &sse.Event{
				//	Data: line,
				//})
			}
		}
	}
}


