package main

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/handlers"
	"github.com/spoonboy-io/lock/internal/help"
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
	raw, err := metadata.GetMetadata(internal.METADATA_URL, logger)
	if err != nil {
		logger.FatalError("problem retrieving metadata", err)
	}

	metadata, err := metadata.ParseMetadataYAML(raw, logger)
	if err != nil {
		logger.FatalError("problem parsing metadata", err)
	}

	//fmt.Println(metadata)
	_ = metadata

	if len(os.Args) < 2 {
		fmt.Printf(help.Options(), version, goversion)
		os.Exit(0)
	}

	args := os.Args[1:]

	switch args[0] {
	case "templates":
		// list all the templates
		var templateInfo string
		templateInfo, err := handlers.ListTemplates(&metadata, args, logger)
		if err != nil {
			logger.FatalError("problem listing templates", err)
		}
		fmt.Println(templateInfo)
		os.Exit(0)
	case "inspect":
		// list tags for a template
		var tagInfo string
		tagInfo, err := handlers.ListTags(logger)
		if err != nil {
			logger.FatalError("problem listing repository tags", err)
		}
		fmt.Println(tagInfo)
		os.Exit(0)
	case "new":
		// create a new project
		if err := handlers.NewProject(args, logger); err != nil {
			logger.FatalError("problem creating new project", err)
		}
	case "watch":
		// monitor, build & watch morpheus
		if err := handlers.Watcher(args); err != nil {
			logger.FatalError("problem watching project", err)
		}
	default:
		// handles help argument also
		fmt.Printf(help.Options(), version, goversion)
		os.Exit(0)
	}
}
