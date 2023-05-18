package gitops

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
	"path/filepath"
)

// GetTags clones the git repository to memory and fetches all the available tags
func GetTags(repo string) ([]string, error) {
	tagList := []string{}

	// clone to memory
	fs := memfs.New()
	storer := memory.NewStorage()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: repo,
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

// CloneRepository clones repository into project folder and checks reference
func CloneRepository(repo, projectName, ref string) error {
	_, err := git.PlainClone(projectName, false, &git.CloneOptions{
		URL:           repo,
		Progress:      nil,
		ReferenceName: plumbing.ReferenceName(ref),
	})
	if err != nil {
		return err
	}

	return nil
}

// DeGit simply removes the .git folder so clean for new git init
func DeGit(projectName string) error {
	gitFolder := filepath.Join(projectName, ".git")
	if err := os.RemoveAll(gitFolder); err != nil {
		return err
	}
	return nil
}
