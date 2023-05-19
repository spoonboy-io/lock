package handlers

import (
	"github.com/spoonboy-io/lock/internal/metadata"
)

func ListPluginVersions(meta *metadata.RssMetadata, args []string) (string, error) {
	return "plugins versions", nil
}
