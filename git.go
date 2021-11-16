package main

import (
	"os"

	"github.com/go-git/go-git/v5"
)

var repos map[string]*git.CloneOptions

func init() {
	repos = map[string]*git.CloneOptions{
		"factorio-init": {
			URL:      "https://github.com/Bisa/factorio-init",
			Progress: os.Stdout,
		},
	}
}

/*
	Clones a git repo into the specified dir.
	repo must be listed in the repos map.
*/
func cloneRepo(repo string, dir string) error {
	_, err := git.PlainClone(dir, false, repos[repo])

	return err
}
