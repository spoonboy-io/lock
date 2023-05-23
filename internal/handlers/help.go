package handlers

// Help returns a multirow string for display on help argument or any unrecognised argument
func Help() string {
	return `
Usage: lock COMMAND <ARGUMENT> --FLAG(S)

A CLI helper tool for Morpheus plugin development and download

General commands:	
  help                   Show this help section
  version                Show the Lock version information

Starter templates:
  templates              List available starter project templates.
                          --category    The category as a filter
                          --morpheus    The minimum version of Morpheus as a filter
  template <TEMPLATE>    Show template information and available versions, <TEMPLATE> can be name or id.
  pick <TEMPLATE>        Creates new project from a starter template, <TEMPLATE> can be name or id.
                          --name        A project folder name (default: %s)
                          --tag         The specific tag to create project from (default: head)

Precompiled plugins:
  plugins                List available Morpheus plugin JAR files.
                          --morpheus    The minimum version of Morpheus as a filter
  plugin <PLUGIN>        Show plugin information and Morpheus version compatibility, <PLUGIN> can be name or id.
  load <PLUGIN>          Download a compiled plugin, <PLUGIN> can be name or id.
                          --version     The specific version of the plugin (default: latest)

`
}
