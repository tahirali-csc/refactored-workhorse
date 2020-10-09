package buildstep

import (
	"net/http"
	"strconv"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
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
