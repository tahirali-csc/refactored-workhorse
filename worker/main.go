package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	ssh2 "golang.org/x/crypto/ssh"
	"log"

	"github.com/workhorse/worker/pkg/executor"
	"github.com/workhorse/worker/pkg/executor/buildstep"
	"net/http"
	"os"
)

func _main() {

	buildOrchestrator := executor.NewBuildOrchestrator()
	go buildOrchestrator.Start()

	stepServer := buildstep.NewServer()

	http.HandleFunc("/runstep", stepServer.HandleRunStep)
	http.HandleFunc("/stream/step", stepServer.HandleLogStream)

	http.ListenAndServe("localhost:8086", nil)

}

func main(){
	err := gitClone("/Users/tahir/workspace/git-tutorial")
	if err != nil {
		log.Println(err)
	}
}

func gitClone(dir string) error{

	publicKeys, err := ssh.NewPublicKeysFromFile("git", "/Users/tahir/.ssh/id_rsa", "")
	publicKeys.HostKeyCallback = ssh2.InsecureIgnoreHostKey()

	if err != nil {
		return err
	}

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: "git@github.com:tahirali-csc/hello-app.git",
		Progress: os.Stdout,
		Auth:     publicKeys,

	})

	if err != nil {
		return err
	}

	return nil
}