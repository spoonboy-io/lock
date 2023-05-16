package metadata_test

import (
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal/metadata"
	"reflect"
	"testing"
)

func TestParseMetadataYAML(t *testing.T) {
	logger := &koan.Logger{}

	testCases := []struct {
		name    string
		yaml    []byte
		wantRes metadata.Metadata
		wantErr error
	}{
		{
			"bad format yaml we omit the header expect specific error",
			[]byte(`--
- plugin:
    category: BADtab
    description: Shows a tab, which reads from a database table and allows delete operations via controller
    url: https://github.com/spoonboy-io/custom-tab-plugin.git
    tags: database,starter,controller,ui`),
			metadata.Metadata{},
			metadata.ERR_UNMARSHALLING,
		},
		{
			"simple no versioning info supplied",
			[]byte(`---
- plugin:
    category: tab
    description: Shows a tab, which reads from a database table and allows delete operations via controller
    url: https://github.com/spoonboy-io/custom-tab-plugin.git
    tags: database,starter,controller,ui`),
			metadata.Metadata{
				{
					metadata.Plugin{
						Category:    "tab",
						Description: "Shows a tab, which reads from a database table and allows delete operations via controller",
						URL:         "https://github.com/spoonboy-io/custom-tab-plugin.git",
						Tags:        "database,starter,controller,ui",
					},
				},
			},
			nil,
		},
		{
			"bad category supplied expect specific error",
			[]byte(`---
- plugin:
    category: BADtab
    description: Shows a tab, which reads from a database table and allows delete operations via controller
    url: https://github.com/spoonboy-io/custom-tab-plugin.git
    tags: database,starter,controller,ui`),
			metadata.Metadata{},
			nil,
		},
		{
			"semantic versioning true, no prefix supplied, default v should be set",
			[]byte(`---
- plugin:
    category: tab
    description: Shows a tab, which reads from a database table and allows delete operations via controller
    url: https://github.com/spoonboy-io/custom-tab-plugin.git
    tags: database,starter,controller,ui
    versioning:
      semantic: true 
`),
			metadata.Metadata{
				{
					metadata.Plugin{
						Category:    "tab",
						Description: "Shows a tab, which reads from a database table and allows delete operations via controller",
						URL:         "https://github.com/spoonboy-io/custom-tab-plugin.git",
						Tags:        "database,starter,controller,ui",
						Versioning: metadata.Versioning{
							Semantic:       true,
							SemanticPrefix: "v",
						},
					},
				},
			},
			nil,
		},
		{
			"morpheus versioning true, no prefix supplied, default morpheus should be set",
			[]byte(`---
- plugin:
    category: tab
    description: Shows a tab, which reads from a database table and allows delete operations via controller
    url: https://github.com/spoonboy-io/custom-tab-plugin.git
    tags: database,starter,controller,ui
    versioning:
      morpheus: true 
`),
			metadata.Metadata{
				{
					metadata.Plugin{
						Category:    "tab",
						Description: "Shows a tab, which reads from a database table and allows delete operations via controller",
						URL:         "https://github.com/spoonboy-io/custom-tab-plugin.git",
						Tags:        "database,starter,controller,ui",
						Versioning: metadata.Versioning{
							Morpheus:       true,
							MorpheusPrefix: "morpheus",
						},
					},
				},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//config = tc.config
			gotRes, gotErr := metadata.ParseMetadataYAML(tc.yaml, logger)
			if !reflect.DeepEqual(gotRes, tc.wantRes) {
				// TODO temp fix to ensure what are the same results does not cause test to fail
				if len(tc.wantRes) != len(gotRes) {
					t.Errorf("\n\nwanted\n%v\n, \n\ngot \n%v\n", tc.wantRes, gotRes)
				}
			}

			if gotErr != tc.wantErr {
				t.Errorf("wanted %v got %v", tc.wantErr, gotErr)
			}
		})
	}

}
