package remotes

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func openRepositoryAndWorktree(path string) (*git.Repository, *git.Worktree, error) {

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, nil, err
	}

	return repo, worktree, nil
}

func CreateNewBranch(path string) error {

	r, _, err := openRepositoryAndWorktree(path)
	if err != nil {
		return err
	}

	if branch, _ := r.Branch("holgersync"); branch != nil {
		return nil
	}

	err = r.CreateBranch(&config.Branch{
		Name: "holgersync",
	})
	if err != nil {
		return err
	}

	return nil
}

func CommitAndPush(path string, targetFile string) error {

	r, w, err := openRepositoryAndWorktree(path)
	if err != nil {
		return err
	}

	w.Add(targetFile)
	w.Commit("hoglersync", &git.CommitOptions{})

	auth := &http.BasicAuth{
		Username: os.Getenv("GITHUB_USERNAME"),
		Password: os.Getenv("GITHUB_PERSONAL_ACCESSTOKEN"),
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		return err
	}

	return nil
}

func CheckFileStatusCode(path string, targetFile string) (*git.StatusCode, error) {

	filePath := filepath.Join(path, targetFile)

	_, w, err := openRepositoryAndWorktree(path)
	if err != nil {
		return nil, err
	}

	ws, err := w.Status()
	if err != nil {
		return nil, err
	}

	status := ws.File(filePath).Worktree

	return &status, nil
}
