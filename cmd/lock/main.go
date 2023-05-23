package main

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/handlers"
	"github.com/spoonboy-io/lock/internal/metadata"
	"os"
)

var (
	version   = "Development build"
	goversion = "Unknown"
)

var logger *koan.Logger

func main() {
	logger = &koan.Logger{}

	// get metadata
	rawYAML, err := metadata.GetMetadata(internal.METADATA_URL, internal.TEMPLATE_CACHE, internal.TEMPLATE_CACHE_TTL, logger)
	if err != nil {
		logger.FatalError("problem retrieving YAML metadata", err)
	}

	rawRSS, err := metadata.GetMetadata(internal.PLUGIN_JAR_INFO_URL, internal.PLUGIN_CACHE, internal.PLUGIN_CACHE_TTL, logger)
	if err != nil {
		logger.FatalError("problem retrieving YAML metadata", err)
	}

	templateMetadata, err := metadata.ParseMetadataYAML(rawYAML, logger)
	if err != nil {
		logger.FatalError("problem parsing template metadata", err)
	}

	pluginMetadata, err := metadata.ParseMetadataXML(rawRSS, logger)
	if err != nil {
		logger.FatalError("problem parsing plugin metadata", err)
	}

	_ = pluginMetadata

	if len(os.Args) < 2 {
		fmt.Printf(handlers.Help(), internal.DEFAULT_PROJECT_NAME)
		os.Exit(0)
	}

	args := os.Args[1:]

	switch args[0] {
	// templates
	case "templates":
		// list all the templates
		var templateInfo string
		templateInfo, err := handlers.ListTemplates(&templateMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(templateInfo)
	case "template":
		// view template metadata and fetch tags for template from remote
		var tagInfo string
		tagInfo, err := handlers.ViewTemplate(&templateMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(tagInfo)
	case "pick":
		// create a new project
		var projectInfo string
		projectInfo, err := handlers.NewProject(&templateMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(projectInfo)
	// jars
	case "plugins":
		// list available plugins
		var pluginsInfo string
		pluginsInfo, err := handlers.ListPlugins(&pluginMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(pluginsInfo)
	case "plugin":
		// view plugin info and version/morpheus reqs
		var pluginInfo string
		pluginInfo, err := handlers.ListPluginVersions(&pluginMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(pluginInfo)
	case "load":
		// download a plugin jar
		var downloadInfo string
		downloadInfo, err := handlers.DownloadPluginVersion(&pluginMetadata, args)
		if err != nil {
			fmt.Printf("%v\n\n", err)
		}
		fmt.Printf(downloadInfo)
	// general
	case "version":
		// handles version command
		fmt.Printf(handlers.Version(), version, goversion, internal.METADATA_URL,
			internal.TEMPLATE_CACHE, (internal.TEMPLATE_CACHE_TTL/1000000000)/60,
			internal.PLUGIN_JAR_INFO_URL, internal.PLUGIN_CACHE,
			(internal.PLUGIN_CACHE_TTL/1000000000)/60, internal.DEFAULT_PROJECT_NAME)
	default:
		// handles help argument also
		fmt.Printf(handlers.Help(), internal.DEFAULT_PROJECT_NAME)
	}
}
