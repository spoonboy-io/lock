package gitops

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spoonboy-io/lock/internal"
)

// ListTags clones the git repository to memory and fetches all the available tags
func ListTags() ([]string, error) {
	tagList := []string{}

	// clone to memory
	fs := memfs.New()
	storer := memory.NewStorage()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: internal.PROJECT_URL,
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
