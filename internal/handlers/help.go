package handlers

// Options returns a multistring for display on help argument or any unrecognised argument
func Help() string {
	return ` 
---------------------------------------------------------                                            
 LOCK - CLI helper tool for Morpheus plugin development
 Version: %s, Go build: %s                    
---------------------------------------------------------

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
