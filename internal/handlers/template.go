package handlers

import (
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/gitops"
	"github.com/spoonboy-io/lock/internal/metadata"
	"strconv"
	"strings"
)

// ViewTemplate provides complete information about a template including the available tags
func ViewTemplate(meta *metadata.Metadata, args []string) (string, error) {
	tagInfo := ""
	output := `Template information:
  ID: %d
  Name: %s
  Description: %s
  Category: %s
  Minimum Morpheus: %s
  Tags: %s
  Repository: %s
%s
`
	var template string
	var p metadata.Plugin

	// check there is id/name argument
	if len(args) < 2 {
		return "", internal.ERR_NO_TEMPLATE
	}
	template = args[1]

	// get template info using either id or key
	id, err := strconv.Atoi(template)
	if err != nil {
		// text, can't convert
		p, id, err = meta.GetTemplateByName(template)
		if err != nil {
			return "", err
		}
	} else {
		// id, converts
		p, err = meta.GetTemplateByIndex(id)
		if err != nil {
			return "", err
		}
	}

	// get repo tags
	tags, err := gitops.GetTags(p.URL)
	if err != nil {
		return tagInfo, err
	}

	// we have tags we need to display them according to the
	// versioning info provided for the repository
	var morpTag, releaseTag, miscTag string
	for _, v := range tags {
		if p.Versioning.Semantic {
			if strings.HasPrefix(v, p.Versioning.SemanticPrefix) {
				// collect the release tags
				// strip the version prefix - it might not be the same repo to repo
				v = strings.TrimPrefix(v, p.Versioning.SemanticPrefix)
				releaseTag += fmt.Sprintf("  %s\n", v)
			}
		}

		if p.Versioning.Morpheus {
			if strings.HasPrefix(v, p.Versioning.MorpheusPrefix) {
				// collect the morpheus tested tags
				// strip the version prefix
				v = strings.TrimPrefix(v, p.Versioning.MorpheusPrefix)
				morpTag += fmt.Sprintf("  %s\n", v)
			}
		}

		if !p.Versioning.Semantic && !p.Versioning.Morpheus {
			// just collect the tags
			miscTag += fmt.Sprintf("  %s\n", v)
		}
	}

	// handle empty despite versioning flags true
	if releaseTag == "" {
		releaseTag = "  No release tags found.\n"
	}

	if morpTag == "" {
		morpTag = "  No Morpheus version tested tags found.\n"
	}

	if miscTag == "" {
		miscTag = "  No tags found.\n"
	}

	// some mini templates to build output
	morpSyntax := `Morpheus version tested tags:
%s`

	relSyntax := `Release tags:
%s`

	miscSyntax := `Tags:
%s`
	// what should output
	tagOutput := ""
	// no versioning just so tags
	if !p.Versioning.Semantic && !p.Versioning.Morpheus {
		tagOutput = fmt.Sprintf(miscSyntax, miscTag)
	}

	// morpheus versioning
	if p.Versioning.Morpheus {
		tagOutput = fmt.Sprintf(morpSyntax, morpTag)
	}

	// release versioning
	if p.Versioning.Semantic {
		tagOutput += fmt.Sprintf(relSyntax, releaseTag)
	}

	url := strings.TrimSuffix(p.URL, ".git")
	tagInfo = fmt.Sprintf(output, id, p.Name, p.Description, p.Category, p.MinimumMorpheus, p.Tags, url, tagOutput)

	return tagInfo, nil
}
