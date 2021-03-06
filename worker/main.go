package main

import (
	"strconv"

	"github.com/workhorse/worker/pkg/executor"
	"github.com/workhorse/worker/pkg/executor/buildstep"
	"net/http"
)

func main() {

	buildOrchestrator := executor.NewBuildOrchestrator()
	go buildOrchestrator.Start()

	stepServer := buildstep.NewServer()

	http.HandleFunc("/runstep", stepServer.HandleRunStep)
	http.HandleFunc("/stream/step", stepServer.HandleLogStream)

	http.HandleFunc("/testbuild", func(writer http.ResponseWriter, request *http.Request) {
		val, _ := request.URL.Query()["buildId"]
		bid, _ := strconv.Atoi(val[0])
		buildOrchestrator.RunBuild(bid)
	})

	http.ListenAndServe("localhost:8086", nil)

}

//func main(){
//	tempDir, err := ioutil.TempDir("", "app")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(tempDir)
//
//	err = executor.GitClone2(tempDir)
//	if err != nil {
//		log.Println(err)
//	}
//}

