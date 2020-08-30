package main

import (
	"net/http"

	"github.com/workhorse/apiserver/pkg/build"
)

func main() {
	buildServer := build.BuildController{}
	//
	http.HandleFunc("/build", buildServer.CreateBuild)
	http.ListenAndServe("localhost:8081", nil)
}
