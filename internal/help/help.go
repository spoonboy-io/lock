package help

// Options returns a multistring for display on help argument or any unrecognised argument
func Options() string {
	return `
LOCK                                            
-----------------------------------------------                                            
CLI helper tool for Morpheus plugin development
Version: %s, Go build: %s                    
-----------------------------------------------

Supported command arguments and flags:

   help                Prints this help section

   templates           Print out a list of starter project templates which are available to select
                       Specify a flag to filter by category:
                         --category    Specify the category as a filter
                         --morpheus	   Specify the minimum version of Morpheus as a filter 

   inspect <template>  Print the git tag references available to choose from for a project template
                       template can be name or id

   new <template>      Creates a new plugin project from a starter template repository,
                       template can be name or id. Specify flags for the template and to override the defaults:
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
