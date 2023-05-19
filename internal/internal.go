package internal

import (
	"errors"
	"time"
)

const (
	METADATA_URL         = "https://raw.githubusercontent.com/spoonboy-io/lock-plugin-metadata/main/lock.yaml"
	TEMPLATE_CACHE       = ".lockTemplateCache"
	TEMPLATE_CACHE_TTL   = 5 * time.Minute
	PLUGIN_JAR_INFO_URL  = "https://share.morpheusdata.com/feed"
	PLUGIN_CACHE         = ".lockPluginCache"
	PLUGIN_CACHE_TTL     = 5 * time.Minute
	DEFAULT_PROJECT_NAME = "morpheus-plugin-project"
)

var (
	ERR_NO_TEMPLATE = errors.New("template id or name not provided")
	ERR_INVALID_TAG = errors.New("requested tag could not be found on remote")
)
