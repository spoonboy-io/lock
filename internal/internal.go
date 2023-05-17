package internal

import (
	"time"
)

const (
	METADATA_URL   = "https://raw.githubusercontent.com/spoonboy-io/lock-plugin-metadata/main/lock.yaml"
	METADATA_CACHE = ".lock_cache"
	CACHE_TTL      = 5 * time.Minute

	PROJECT_URL          = "https://github.com/spoonboy-io/switch.git"
	DEFAULT_PROJECT_NAME = "morpheus-plugin-project"
)
