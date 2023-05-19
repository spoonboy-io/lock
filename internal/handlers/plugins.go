package handlers

import (
	"github.com/spoonboy-io/lock/internal/metadata"
)

func ListPlugins(meta *metadata.RssMetadata, args []string) (string, error) {

	return "plugins", nil
}
