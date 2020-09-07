package main

import (
	"log"
	"net/http"

	"github.com/workhorse/apiserver/pkg/build"
	"github.com/workhorse/apiserver/pkg/config"
)

func main() {
	config := config.GetAppConfig()
	log.Println(*config)

	buildServer := build.BuildController{}
	//
	http.HandleFunc("/build", buildServer.CreateBuild)
	http.HandleFunc("/buildbinding", buildServer.BindToNode)
	http.HandleFunc("/updatestepstatus", buildServer.UpdateBuildStepStatus)
	http.HandleFunc("/updatebuildstepbinding", buildServer.BindingBuildStepToNode)

	http.ListenAndServe("localhost:"+config.Server.Port, nil)
}
