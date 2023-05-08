package handlers

import (
	"errors"
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"strings"
)

// NewProject handles the creation of a new plugin project folder cloned
// from either head or a specfic git tag of the template-plugin github project
func NewProject(args []string) error {
	// set defaults
	projectName := internal.DEFAULT_PROJECT_NAME
	reference := ""

	// parse flags, we don't use flags package as requires flags
	// precede the arguments which don't want, so we do ourselves
	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--name=") {
			projectName = strings.TrimLeft(v, "--name=")
		}
		if strings.HasPrefix(v, "--tag=") {
			reference = strings.TrimLeft(v, "--tag=")
		}
	}

	// valid tag (if not default)
	if reference != "" {
		tags, err := internal.ListTags()
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

	// tag is good so clone the repo and checkout that ref (or head)
	fmt.Println(projectName, reference)

	return nil
}

/*


 */

