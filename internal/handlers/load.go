package handlers

import (
	"github.com/spoonboy-io/lock/internal/metadata"
)

func DownloadPluginVersion(meta *metadata.RssMetadata, args []string) (string, error) {
	return "download", nil
}
