package storage

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type GitStorage struct {
	LocalPath   string `yaml:"local_path"`
	jsonStorage *JsonStorage
}

func (gs *GitStorage) initJsonStorage() {
	gs.jsonStorage = &JsonStorage{Path: gs.LocalPath}
}

func (gs *GitStorage) Sync() error {
	gs.initJsonStorage()

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
	if err != nil {
		fmt.Println("encountered issue trying to sync git repo")
		return fmt.Errorf("encountered an error trying to pull from remote origin: %w", err)
	}
	return nil
}

func (gs *GitStorage) GetLinkForLabel(label string) string {
	gs.initJsonStorage()

	err := gs.Sync()
	if err != nil {
		return ""
	}
	return gs.jsonStorage.GetLinkForLabel(label)
}

func (gs *GitStorage) GetAllLabels() []string {
	gs.initJsonStorage()

	err := gs.Sync()
	if err != nil {
		return []string{}
	}
	return gs.jsonStorage.GetAllLabels()
}

func (gs *GitStorage) AddLink(label string, url string) error {
	gs.initJsonStorage()

	err := gs.Sync()
	if err != nil {
		return err
	}
	err = gs.jsonStorage.AddLink(label, url)
	if err != nil {
		return err
	}
	// TODO: COMMIT AND PUSH CHANGES
	return nil
}

func (gs *GitStorage) RemoveLink(label string) error {
	gs.initJsonStorage()

	err := gs.Sync()
	if err != nil {
		return err
	}
	err = gs.jsonStorage.RemoveLink(label)
	if err != nil {
		return err
	}
	// TODO: COMMIT AND PUSH CHANGES
	return nil
}
