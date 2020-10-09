package main

import (
	"github.com/workhorse/worker/pkg/executor"
	"github.com/workhorse/worker/pkg/executor/buildstep"
	"net/http"
)

func main() {

	buildOrchestrator := executor.NewBuildOrchestrator()
	go buildOrchestrator.Start()

	stepServer := buildstep.NewServer()

	http.HandleFunc("/runstep", stepServer.HandleRunStep)

	http.ListenAndServe("localhost:8086", nil)

}
