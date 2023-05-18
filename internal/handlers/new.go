package handlers

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/gitops"
	"github.com/spoonboy-io/lock/internal/metadata"
	"os"
	"strconv"
	"strings"
)

// NewProject handles the creation of a new starter plugin project cloned
// from either head or a specific git tag of the template repository
func NewProject(meta *metadata.Metadata, args []string, logger *koan.Logger) (string, error) {

	// set defaults
	projectName := internal.DEFAULT_PROJECT_NAME
	reference := ""

	// parse flags, we don't use flags package as requires flags
	// precede the arguments which don't want, so we do ourselves
	// we collect the id
	var template, fullTag string
	var p metadata.Plugin

	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--name=") {
			projectName = strings.TrimPrefix(v, "--name=")
		}
		if strings.HasPrefix(v, "--tag=") {
			reference = strings.TrimPrefix(v, "--tag=")
		}
		// anything else should be the template?
		if !strings.HasPrefix(v, "--name=") && !strings.HasPrefix(v, "--tag=") {
			template = v
		}
	}

	// if empty we have an problem
	if template == "" {
		return "", internal.ERR_NO_TEMPLATE
	}

	// get template info using either id or key
	id, err := strconv.Atoi(template)
	if err != nil {
		// text, can't convert
		p, id, err = meta.GetByName(template)
		if err != nil {
			return "", err
		}
	} else {
		// id, converts
		p, err = meta.GetByIndex(id)
		if err != nil {
			return "", err
		}
	}

	// valid tag (if not default)
	if reference != "" {
		//logger.Info("checking tag exists")
		tags, err := gitops.GetTags(p.URL)
		if err != nil {
			return "", err
		}

		validRef := false

		// TODO the way this is done leaves a very small risk of collisions
		// user specifies --tag=6.0.1 which relates to `morpheus6.0.1` and the plugin also has a semantic version
		// at the same version number `v6.0.1`. Tiny but worth noting
		for _, v := range tags {
			// check the tag
			if p.Versioning.Semantic {
				if v == fmt.Sprintf("%s%s", p.Versioning.SemanticPrefix, reference) {
					validRef = true
					fullTag = v
				}
			}

			if p.Versioning.Morpheus {
				if v == fmt.Sprintf("%s%s", p.Versioning.MorpheusPrefix, reference) {
					validRef = true
					fullTag = v
				}
			}

			if !p.Versioning.Semantic && !p.Versioning.Morpheus {
				if v == reference {
					validRef = true
					fullTag = v
				}
			}

		}

		if !validRef {
			return "", internal.ERR_INVALID_TAG
		}
	}

	// create the project folder
	if err := os.Mkdir(projectName, 0755); err != nil {
		return "", err
	}

	// first set up reference to clone if a tag
	fullRef := ""
	if reference != "" {
		fullRef = fmt.Sprintf("refs/tags/%s", fullTag)
	}

	// clone the repository into the directory
	if err := gitops.CloneRepository(p.URL, projectName, fullRef); err != nil {
		return "", err
	}

	// remove .git folder
	if err := gitops.DeGit(projectName); err != nil {
		return "", err
	}

	output := `
New Plugin starter project successfully created
-----------------------------------------------
Created in folder: '%s'
Using plugin template: '%s'
Using tag/reference: '%s'.

`
	if reference == "" {
		reference = "head"
	}
	return fmt.Sprintf(output, projectName, p.Name, reference), nil
}
