package storage

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

const fileName string = "links.json"

type GitStorage struct {
	LocalPath   string `yaml:"local_path"`
	AutoSync    bool   `yaml:"auto_sync"`
	jsonStorage *JsonStorage
}

func (gs *GitStorage) initJsonStorage() {
	gs.jsonStorage = &JsonStorage{Path: fmt.Sprintf("%s%s", gs.LocalPath, fileName)}
}

func generateCommitMessage(label string, isAdd bool) string {
	if isAdd {
		return fmt.Sprintf("Adding '%s'", label)
	} else {
		return fmt.Sprintf("Removing '%s'", label)
	}
}

func (gs *GitStorage) commitAndPush(label string, isAdd bool) error {
	repo, err := git.PlainOpen(gs.LocalPath)
	if err != nil {
		return fmt.Errorf("encountered an error trying to open local git repo: %w", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("encountered an error trying to get the working tree: %w", err)
	}

	_, err = w.Add(fileName)
	if err != nil {
		return fmt.Errorf("encountered error trying to add '%s': %w", fileName, err)
	}

	fmt.Println("committing changes...")
	_, err = w.Commit(createCommitMessage(label, isAdd), &git.CommitOptions{All: true})
	if err != nil {
		return fmt.Errorf("encountered an error trying to commit changes: %w", err)
	}
	fmt.Println("pushing changes...")
	err = repo.Push(&git.PushOptions{})
	if err != nil {
		return fmt.Errorf("encounterd an error trying to push changes: %w", err)
	}
	fmt.Println("done")
	return nil
}

func (gs *GitStorage) Sync() error {
	fmt.Println("syncing...")
	repo, err := git.PlainOpen(gs.LocalPath)
	if err != nil {
		return fmt.Errorf("encountered an error trying to open local git repo: %w", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("encountered an error trying to get the working tree: %w", err)
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("encountered an error trying to pull from remote origin: %w", err)
	}
	fmt.Println("done syncing")
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

	err := gs.jsonStorage.AddLink(label, url)
	if err != nil {
		return err
	}

	if err := gs.commitAndPush(label, true); err != nil {
		fmt.Println(err)
	}
	return err
}

func (gs *GitStorage) RemoveLink(label string) error {
	if gs.AutoSync {
		if err := gs.Sync(); err != nil {
			fmt.Println(err)
		}
	}

	gs.initJsonStorage()

	err := gs.jsonStorage.RemoveLink(label)
	if err != nil {
		return fmt.Errorf("error trying to remove link from local file: %w", err)
	}

	if err := gs.commitAndPush(label, false); err != nil {
		return err
	}
	return nil
}
