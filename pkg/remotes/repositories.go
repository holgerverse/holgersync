package remotes

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CreateNewBranch(path string) error {

	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	if branch, _ := repo.Branch("holgersync"); branch != nil {
		return nil
	}

	err = repo.CreateBranch(&config.Branch{
		Name: "holgersync",
	})
	if err != nil {
		return err
	}

	return nil
}

func CommitAndPush(path string, targetFile string) error {

	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, _ := repo.Worktree()
	w.Add(targetFile)
	w.Commit("hoglersync", &git.CommitOptions{})

	auth := &http.BasicAuth{
		Username: os.Getenv("GITHUB_USERNAME"),
		Password: os.Getenv("GITHUB_PERSONAL_ACCESSTOKEN"),
	}

	fmt.Println(auth)

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		return err
	}

	return nil

}
