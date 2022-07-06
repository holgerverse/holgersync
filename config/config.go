package config

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/spf13/viper"
)

type Config struct {
	HolgersyncConfig HolgersyncConfig
	Logger           Logger
}

type HolgersyncConfig struct {
	SourceFileConfig SourceFileConfig `yaml:"sourceFileConfig"`
	Targets          []Target         `mapstructure:"Targets"`
}

type GlobalConfig struct {
	Git []GitConfig `mapstructure:"gitConfig"`
}

type GitConfig struct {
	Username            string `yaml:"username"`
	PersonalAccessToken string `yaml:"personalAccessToken"`
	Remote              string `yaml:"remote"`
	Branch              string `yaml:"branch"`
}

type SourceFileConfig struct {
	FilePath  string `yaml:"filePath"`
	RootPath  string `yaml:"rootPath"`
	FileRegex string `yaml:"fileRegex"`
}

type Target struct {
	Path       string      `yaml:"path"`
	Git        []GitConfig `mapstructure:"gitConfig"`
	Parameters []Parameter `mapstructure:"parameters"`
}

type Parameter struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Logger struct {
	Level       string
	Destination string
}

// Load the holgersync config file
func LoadConfig(filename string) (*viper.Viper, error) {

	v := viper.New()

	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// TODO
func ParseConfig(v *viper.Viper) *Config {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &c
}

func (t *Target) openRepositoryAndWorktree() (*git.Repository, *git.Worktree, error) {

	repo, err := git.PlainOpen(t.Path)
	if err != nil {
		return nil, nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, nil, err
	}

	return repo, worktree, nil
}

// Create the desired Holgersync branch, if the branch already exists do nothing
func (t *Target) CreateHolgersyncBranch() error {

	// Variable to store the branch name, will be "holgersync" if not specified
	var b string = "holgersync"

	// Open the repostiroy and worktree
	r, _, err := t.openRepositoryAndWorktree()
	if err != nil {
		return err
	}

	// Create reference for hoglersync branch
	branchName := plumbing.NewBranchReferenceName(b)

	// Get the reference to the HEAD of the repository
	headRef, err := r.Head()
	if err != nil {
		return err
	}

	// Create refrence for the new branch pointing to the HEAD of the repository
	ref := plumbing.NewHashReference(branchName, headRef.Hash())

	// Create the new branch and write it to the repository
	err = r.Storer.SetReference(ref)
	if err != nil {
		return fmt.Errorf("failed to create branch %s: %s", b, err)
	}

	return nil
}
