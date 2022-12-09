package storage

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

type GitStorage struct {
	LocalPath   string `yaml:"local_path"`
	AutoSync    bool   `yaml:"auto_sync,omitempty"`
	jsonStorage *JsonStorage
}

func (gs *GitStorage) initJsonStorage() {
	gs.jsonStorage = &JsonStorage{Path: gs.LocalPath}
}

func (gs *GitStorage) Sync() error {
	repo, err := git.PlainOpen(gs.LocalPath)
	if err != nil {
		fmt.Println("encountered issue trying to sync git repo")
		return fmt.Errorf("encountered an error trying to open local git repo: %w", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Println("encountered issue trying to sync git repo")
		return fmt.Errorf("encountered an error trying to get the working tree: %w", err)
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		fmt.Println("encountered issue trying to sync git repo")
		return fmt.Errorf("encountered an error trying to pull from remote origin: %w", err)
	}
	return nil
}

func (gs *GitStorage) GetLinkForLabel(label string) string {
	if gs.AutoSync {
		if err := gs.Sync(); err != nil {
			fmt.Println(err)
		}
	}

	gs.initJsonStorage()

	return gs.jsonStorage.GetLinkForLabel(label)
}

func (gs *GitStorage) GetAllLabels() []string {
	if gs.AutoSync {
		if err := gs.Sync(); err != nil {
			fmt.Println(err)
		}
	}

	gs.initJsonStorage()

	return gs.jsonStorage.GetAllLabels()
}

func (gs *GitStorage) AddLink(label string, url string) error {
	if gs.AutoSync {
		if err := gs.Sync(); err != nil {
			fmt.Println(err)
		}
	}

	gs.initJsonStorage()

	return gs.jsonStorage.AddLink(label, url)
}

func (gs *GitStorage) RemoveLink(label string) error {
	if gs.AutoSync {
		if err := gs.Sync(); err != nil {
			fmt.Println(err)
		}
	}

	gs.initJsonStorage()

	return gs.jsonStorage.RemoveLink(label)
}
