package metadata

import (
	"errors"
	"github.com/spoonboy-io/koan"
	"gopkg.in/yaml.v3"
	"strings"
)

// Metadata is a representation of the parsed YAML metadata
type Metadata []struct {
	Plugin `yaml:"plugin"`
}

// Plugin represents the metadata for a single plugin starter
type Plugin struct {
	Name            string
	Category        string     `yaml:"category"`
	Description     string     `yaml:"description"`
	URL             string     `yaml:"url"`
	Tags            string     `yaml:"tags"`
	Versioning      Versioning `yaml:"versioning"`
	MinimumMorpheus string     `yaml:"minimumMorpheus"`
	Disabled        bool       `yaml:"disabled"`
}

// Versioning represents metadata which describes versioning semantics
type Versioning struct {
	Semantic       bool   `yaml:"semantic"`
	SemanticPrefix string `yaml:"semanticPrefix"`
	Morpheus       bool   `yaml:"morpheus"`
	MorpheusPrefix string `yaml:"morpheusPrefix"`
}

const (
	// defaults
	SEMANTIC_PREFIX = "v"
	MORPHEUS_PREFIX = "morpheus"
)

// curently we perform category checks to make sure it is one of those
// which represents a morpheus plugin provider, this will be a workflow action
// in the metadata repository eventually
var category = [9]string{
	"approval",
	"cypher",
	"dns",
	"tab",
	"ipam",
	"task",
	"report",
	"cloud",
	"backup",
}

var ERR_UNMARSHALLING = errors.New("could not unmarshal the YAML")

// ParseMetadataYAML parses the YAML metadata to the Metadata type
// then sets defaults on string types which would otherwise be default
// empty, we also check the category is one of those allowed
func ParseMetadataYAML(data []byte, logger *koan.Logger) (Metadata, error) {
	var temp, meta Metadata

	if err := yaml.Unmarshal(data, &temp); err != nil {
		logger.Error("unmarshalling YAML", err)
		return meta, ERR_UNMARSHALLING
	}

	// config checks
	for i, p := range temp {
		// set defaults if needed not set
		if p.Versioning.Semantic {
			if p.Versioning.SemanticPrefix == "" {
				temp[i].Versioning.SemanticPrefix = SEMANTIC_PREFIX
			}
		}
		if p.Versioning.Morpheus {
			if p.Versioning.MorpheusPrefix == "" {
				temp[i].Versioning.MorpheusPrefix = MORPHEUS_PREFIX
			}
		}

		// check category is allowed, and and not disabled and append to meta
		if ok := validCategory(p.Category); ok {
			if !p.Disabled {
				// parse the name
				temp[i].Name = stripName(p.URL)
				meta = append(meta, temp[i])
			}
		}
	}

	return meta, nil
}

func validCategory(cat string) bool {
	var valid bool
	for _, c := range category {
		if c == cat {
			valid = true
		}
	}
	return valid
}

func stripName(uri string) string {
	temp := strings.Split(uri, "/")
	temp2 := temp[len(temp)-1:]
	return strings.TrimSuffix(temp2[0], ".gitops")
}
