package handlers

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal/gitops"
	"strings"
)

// ListTags formats a slice of tags into meaningful output. It uses internal.ListTags
// to perform the git operations and generate the slice of tags which are parsed here
func ListTags(logger *koan.Logger) (string, error) {
	tagInfo := ""
	template := `

Morpheus version tested tags
----------------------------
These plugin-template versions have been tested against the specific version of Morpheus 
mentioned in the tag:

%s
Development release tags
------------------------
These tags are release versions of the plugin-template:

%s
`
	logger.Info("fetching tags from repository")
	var morp, release string
	tags, err := gitops.ListTags()
	if err != nil {
		return tagInfo, err
	}

	for _, v := range tags {
		if strings.HasPrefix(v, "v") {
			// release
			release += fmt.Sprintf("%s\n", v)
		}
		if strings.HasPrefix(v, "morpheus") {
			// morpheus tested
			morp += fmt.Sprintf("%s\n", v)
		}
	}

	if release == "" {
		release = "- No release tags found.\n"
	}

	if morp == "" {
		morp = "- No Morpheus version tested tags found.\n"
	}

	tagInfo = fmt.Sprintf(template, morp, release)

	return tagInfo, nil
}
