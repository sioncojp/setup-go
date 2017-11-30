package main

import (
	"log"
	"os"
	"strings"
)

const (
	app = "go-setup"
)

// reponame ...RepositoryName
type reponame string

// Templates ...plan to create files from go-template
type Templates struct{}

// Dirs ...plan to create diretctorys
type Dirs struct{}

// Files ...plan to create empty files
type Files struct{}

// Material ...specifies file, directory for setup
type Material struct {
	Templates Templates
	Dirs      Dirs
	Files     Files
}

// Option ...operation for setup
type Option interface {
	Get()
	Create()
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix(app + ": ")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run ...run go-setup
func Run() error {
	rn, err := NewRepoName()
	if err != nil {
		return err
	}

	material := &Material{}
	material.Dirs.Create(rn)
	material.Files.Create(rn)
	material.Templates.Create(rn)

	return nil
}

// RepoName ...Return current repository name
func NewRepoName() (reponame, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	s := strings.Split(pwd, "/")

	return reponame(s[len(s)-1]), nil
}

// FileExists ...check file exist
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
