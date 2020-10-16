package main

import (
	"github.com/workhorse/apiserver/pkg/build"
	"github.com/workhorse/apiserver/pkg/config"
	"github.com/workhorse/apiserver/pkg/node"
	"log"
	"net/http"
)

func main() {

	config := config.GetAppConfig()
	log.Println(*config)

	buildServer := build.BuildController{}
	nodeServer := node.NodeInfoServer{}

	//
	http.HandleFunc("/build", buildServer.CreateBuild)
	http.HandleFunc("/buildbinding", buildServer.BindToNode)
	http.HandleFunc("/updatestepstatus", buildServer.UpdateBuildStepStatus)
	http.HandleFunc("/updatebuildstepbinding", buildServer.BindingBuildStepToNode)

	http.HandleFunc("/getbuild", buildServer.GetBuild)
	http.HandleFunc("/getstep", buildServer.GetStep)
	http.HandleFunc("/patch", buildServer.Patch)

	http.HandleFunc("/api/nodeinfo", nodeServer.Handle)

	http.ListenAndServe("localhost:"+config.Server.Port, nil)
}
