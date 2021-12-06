package main

import (
	"github.com/mattermost/cicd-sdk/pkg/cherrypicker"
	"github.com/sirupsen/logrus"
)

func main() {
	/*
		sm := secret.NewManager()
		secret, err := sm.Get("kv/data/a")
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Printf("Salio: %s", secret.Value.GetString("test"))
	*/

	cp := cherrypicker.NewWithOptions(&cherrypicker.Options{
		// RepoPath:  "/home/urbano/Projects/Mattermost/mattermost-server",
		RepoOwner: "mattermost",
		RepoName:  "mattermost-webapp",
		ForkOwner: "puerco",
	})

	//if err := cp.CreateCherryPickPR(9340, "release-6.1"); err != nil {
	if err := cp.CreateCherryPickPR(9410, "release-6.2"); err != nil {
		logrus.Error(err.Error())
	}
}
