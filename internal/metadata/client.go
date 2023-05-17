package metadata

import (
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"io"
	"net/http"
	"os"
	"time"
)

// GetMetadata wraps a http request to obtain metadata from the
// github repository in which it is maintained, or from a local cache
// since this is a CLI we don't want to make the same HTTP request on
// every invocation so we cache the data with a TTL
func GetMetadata(uri string) ([]byte, error) {
	var data []byte

	// check cached exist
	if haveCached(internal.METADATA_CACHE, internal.CACHE_TTL) {
		data, err := os.ReadFile(internal.METADATA_CACHE)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: '%d'", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// create cache
	if err := os.WriteFile(internal.METADATA_CACHE, data, 0700); err != nil {
		return nil, err
	}

	return data, nil
}

func haveCached(cache string, ttl time.Duration) bool {
	// cache file exist
	c, err := os.Stat(cache)
	if err != nil {
		return false
	}

	// not expired
	ts := c.ModTime()
	tsAge := time.Now().Sub(ts)
	if tsAge > ttl {
		return false
	}

	return true
}
