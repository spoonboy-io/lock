package handlers

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal/gitops"
	"github.com/spoonboy-io/lock/internal/metadata"
	"strings"
)

// Inspect provides complete information about a template including the available tags
func Inspect(meta *metadata.Metadata, args []string, logger *koan.Logger) (string, error) {
	tagInfo := ""
	template := `
%s
----------------------------
%s

- Category: %s
- Minimum Morpheus: %s
- Tags: %s
- Repository: %s

Morpheus version tested tags
----------------------------

%s
Release tags
------------

%s
`
	p, err := meta.GetByIndex(1)
	if err != nil {
		fmt.Println(err)
	}

	p, err = meta.GetByName(p.Name)
	if err != nil {
		fmt.Println(err)
	}

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

	name := strings.ToUpper(strings.Replace(p.Name, "-", " ", -1))
	tagInfo = fmt.Sprintf(template, name, p.Description, p.Category, p.MinimumMorpheus, p.Tags, p.URL, morp, release)

	return tagInfo, nil
}
