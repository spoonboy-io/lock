package internal

import (
	"errors"
	"fmt"
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

	ERR_NO_PLUGIN = errors.New("plugin id or name not provided")
)

// CutString will cut and suffix a string at specified length
func CutString(data string, cutAt int) string {
	d := []rune(data)
	short := ""
	if len(d) > cutAt {
		if string(d[cutAt-1]) == " " {
			cutAt--
		}
		short = fmt.Sprintf("%s..", string(d[0:cutAt]))
	}
	return short
}

// WriteHeader will write the header footer to correct length
func Writeline(num int) string {
	l := ""
	for i := 0; i < num; i++ {
		l += "-"
	}
	return l
}

// WriteTitle will write the tile and pad correctly
func WriteTitle(key string, num int) string {
	t := key
	c := num - len(t)
	for i := 0; i < c; i++ {
		t += " "
	}
	return t
}

// WriteCell will space a key correctly given field width
func WriteCell(key string, num int) string {
	for i := len(key); i < num; i++ {
		key += " "
	}
	return key
}
