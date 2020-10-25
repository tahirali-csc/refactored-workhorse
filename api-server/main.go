package main

import (
	"github.com/workhorse/apiserver/pkg/build"
	"github.com/workhorse/apiserver/pkg/config"
	"github.com/workhorse/apiserver/pkg/node"
	"github.com/workhorse/apiserver/pkg/project"
	"log"
	"net/http"
)

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {

	config := config.GetAppConfig()
	log.Println(*config)

	buildServer := build.BuildController{}
	nodeServer := node.NodeInfoServer{}
	projectServer := project.NewProjectServer()

	//
	//http.HandleFunc("/build", buildServer.CreateBuild)
	//http.HandleFunc("/buildbinding", buildServer.BindToNode)
	//http.HandleFunc("/updatestepstatus", buildServer.UpdateBuildStepStatus)
	//http.HandleFunc("/updatebuildstepbinding", buildServer.BindingBuildStepToNode)
	//http.HandleFunc("/streamstep/tailog", buildServer.TailLogStep)
	//
	////http.HandleFunc("/getbuild", buildServer.GetBuild)
	//http.HandleFunc("/getstep", buildServer.GetStep)
	//http.HandleFunc("/patch", buildServer.Patch)
	//
	//http.HandleFunc("/api/nodeinfo", nodeServer.Handle)


	mux := http.NewServeMux()
	mux.HandleFunc("/api/project", projectServer.Handle)
	mux.HandleFunc("/api/build", buildServer.Handle)
	mux.HandleFunc("/api/buildsteps", buildServer.HandleBuildStep)

	mux.HandleFunc("/buildbinding", buildServer.BindToNode)
	mux.HandleFunc("/updatestepstatus", buildServer.UpdateBuildStepStatus)
	mux.HandleFunc("/updatebuildstepbinding", buildServer.BindingBuildStepToNode)
	mux.HandleFunc("/streamstep/tailog", buildServer.TailLogStep)

	//http.HandleFunc("/getbuild", buildServer.GetBuild)
	mux.HandleFunc("/getstep", buildServer.GetStep)
	mux.HandleFunc("/patch", buildServer.Patch)

	mux.HandleFunc("/api/nodeinfo", nodeServer.Handle)


	http.ListenAndServe("localhost:"+config.Server.Port, setHeaders(mux))
}
