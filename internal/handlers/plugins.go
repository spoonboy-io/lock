package handlers

import (
	"github.com/spoonboy-io/lock/internal/metadata"
)

// ListPlugins will process metadata and provide output
// related to the available starter plugin jar versions
func ListPlugins(meta *metadata.RssMetadata, args []string) (string, error) {

	return "plugins", nil
}
