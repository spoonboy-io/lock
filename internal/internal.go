package internal

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

const (
	//PROJECT_URL          = "https://github.com/spoonboy-io/plugin-template.git"
	PROJECT_URL          = "https://github.com/spoonboy-io/switch.git"
	DEFAULT_PROJECT_NAME = "morpheus-plugin-project"
)

func Help() string {
	return `

Lock usage information:
-----------------------
Lock should be run with a single argument suffix i.e. lock <argument>
Supported arguments are:

  help   Prints this help section

  tags   Print the git tag references available to choose from.
         Two types of tag exist, one tagged and tested against Morpheus version,
         and second which represents a semantic release.

  new    Creates a new module project from a template repository,
         defaults to folder 'morpheus-plugin-project' and 'head'.
         Specify flags to override the defaults:
           --name   Specify a project folder name
           --tag    Specify a tag to create project from

  watch  Starts a watcher which will build the plugin on save of files and upload
         it to Morpheus, while monitoring Morpheus errors.
         Requires flags to authenticate with Morpheus over REST API:
           --host   The url of the Morpheus appliance
           --token  A bearer token with which to authenticate (exclude the BEARER prefix)`
}

func ListTags() ([]string, error) {
	tagList := []string{}

	url := PROJECT_URL

	// clone to memory
	fs := memfs.New()
	storer := memory.NewStorage()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return tagList, err
	}

	tags, err := r.Tags()
	if err != nil {
		return tagList, err
	}

	err = tags.ForEach(
		func(ref *plumbing.Reference) error {
			obj, err := r.TagObject(ref.Hash())
			if err != nil {
				return err
			}
			tagList = append(tagList, obj.Name)
			return nil
		})
	if err != nil {
		return tagList, err
	}
	return tagList, nil
}
