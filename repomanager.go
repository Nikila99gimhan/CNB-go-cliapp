package main

import (
	"os"
	"os/exec"
	"strings"
)

type RepoManager struct {}

func NewRepoManager() *RepoManager {
	return &RepoManager{}
}

func (rm *RepoManager) GetRepoName(repoURL string) string {
	parts := strings.Split(repoURL, "/")
	return strings.TrimSuffix(parts[len(parts)-1], ".git")
}

func (rm *RepoManager) RepoExists(directory string) bool {
	_, err := os.Stat(directory)
	return !os.IsNotExist(err)
}

func (rm *RepoManager) CloneRepo(repoURL string) error {
	cmd := exec.Command("git", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
