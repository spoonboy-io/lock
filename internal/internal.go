package internal

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"time"
)

const (
	METADATA_URL   = "https://raw.githubusercontent.com/spoonboy-io/lock-plugin-metadata/main/lock.yaml"
	METADATA_CACHE = ".lock_cache"
	CACHE_TTL      = 5 * time.Minute

	PROJECT_URL          = "https://github.com/spoonboy-io/switch.git"
	DEFAULT_PROJECT_NAME = "morpheus-plugin-project"
)

var ()

// Help returns a multistring for display on help argument or any unrecognised argument
func Help() string {
	return `
 LOCK                                            
 -----------------------------------------------                                            
 CLI helper tool for Morpheus plugin development
 Version: %s, Go build: %s                    
 -----------------------------------------------

 Lock should be run with a single argument suffix i.e. lock <argument>
 Supported arguments are:

   help         Prints this help section

   templates    Print out a list of starter project templates which are available to select
                Specify a flag to filter by category:
                  --category    Specify the category as a filter
				  --morpheus	Specify the minimum version of Morpheus as a filter 

   inspect      Print the git tag references available to choose from for a project template
                Specify flag for the template:
                  --template    Specify the template to inspect (required)

   new          Creates a new plugin project from a starter template repository,
                Specify flags for the template and to override the defaults:
                  --template    Specify the template use (required) 
                  --name        Specify a project folder name (default: morpheus-plugin)
                  --tag         Specify a tag to create project from (default: head)
`
}

/*


  watch  Starts a watcher which will build the plugin on save of files and upload
         it to Morpheus, while monitoring Morpheus errors.
         Requires flags to authenticate with Morpheus over REST API:
           --host   The url of the Morpheus appliance
           --token  A bearer token with which to authenticate (exclude the BEARER prefix)
*/

// ListTags clones the git repository to memory and fetches all the available tags
func ListTags() ([]string, error) {
	tagList := []string{}

	// clone to memory
	fs := memfs.New()
	storer := memory.NewStorage()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: PROJECT_URL,
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
