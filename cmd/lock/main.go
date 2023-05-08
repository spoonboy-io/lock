package main

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/handlers"
	"os"
)

var logger *koan.Logger

func printHelpExit() {
	fmt.Println(internal.Help())
	os.Exit(0)
}

func main() {
	logger = &koan.Logger{}

	if len(os.Args) < 2 {
		printHelpExit()
	}

	args := os.Args[1:]

	switch args[0] {
	case "tags":
		// list tags
		if tagInfo, err := handlers.ListTags(logger); err != nil {
			logger.FatalError("problem listing repository tags", err)
		} else {
			fmt.Println(tagInfo)
			os.Exit(0)
		}
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
		printHelpExit()
	}
}
