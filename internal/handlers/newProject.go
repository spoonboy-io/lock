package handlers

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/gitops"
	"os"
	"path/filepath"
	"strings"
)

// NewProject handles the creation of a new plugin project folder cloned
// from either head or a specfic gitops tag of the template-plugin github project
func NewProject(args []string, logger *koan.Logger) error {
	// set defaults
	projectName := internal.DEFAULT_PROJECT_NAME
	reference := ""

	// parse flags, we don't use flags package as requires flags
	// precede the arguments which don't want, so we do ourselves
	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--name=") {
			projectName = strings.TrimPrefix(v, "--name=")
		}
		if strings.HasPrefix(v, "--tag=") {
			reference = strings.TrimPrefix(v, "--tag=")
		}
	}

	// valid tag (if not default)
	if reference != "" {
		logger.Info("checking tag exists")
		tags, err := gitops.ListTags()
		if err != nil {
			return err
		}

		validRef := false
		for _, v := range tags {
			if v == reference {
				validRef = true
			}
		}
		if !validRef {
			return errors.New("not a valid tag")
		}
	}

	// create the project folder
	logger.Info(fmt.Sprintf("creating new project folder '%s'", projectName))
	if err := os.Mkdir(projectName, 0755); err != nil {
		return err
	}

	// clone the repository into the directory
	// logging convenience
	target := fmt.Sprintf("tag '%s'", reference)
	if reference == "" {
		target = "'latest'"
	}
	logger.Info(fmt.Sprintf("cloning repository, checking out %s", target))

	// first set up reference to clone if a tag
	if reference != "" {
		reference = fmt.Sprintf("refs/tags/%s", reference)
	}

	if _, err := git.PlainClone(projectName, false, &git.CloneOptions{
		URL:           internal.PROJECT_URL,
		Progress:      nil,
		ReferenceName: plumbing.ReferenceName(reference),
	}); err != nil {
		return err
	}

	// remove .gitops folder, so clean for new gitops project
	gitFolder := filepath.Join(projectName, ".gitops")
	if err := os.RemoveAll(gitFolder); err != nil {
		return err
	}

	logger.Info("clone completed, new project folder is ready.")
	return nil
}
