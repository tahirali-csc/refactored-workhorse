package executor

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/workhorse/api"
	ssh2 "golang.org/x/crypto/ssh"
	"os"
)

func GitClone(cloneDir string, project api.Project) error {

	//publicKeys, err := ssh.NewPublicKeysFromFile("git", "/Users/tahir/.ssh/id_rsa", "")
	publicKeys, err := ssh.NewPublicKeys("git", []byte(project.PrivateKey), "")
	publicKeys.HostKeyCallback = ssh2.InsecureIgnoreHostKey()

	if err != nil {
		return err
	}

	_, err = git.PlainClone(cloneDir, false, &git.CloneOptions{
		//URL:      "git@github.com:tahirali-csc/hello-app.git",
		URL:      project.CloneURL,
		Progress: os.Stdout,
		Auth:     publicKeys,
		Depth:    1,
	})

	if err != nil {
		return err
	}

	return nil
}
