package metadata

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"io"
	"net/http"
	"os"
	"time"
)

// GetMetadata wraps a http request to obtain metadata from the or from a local cache, since this
// is a CLI we don't want to make the same HTTP request on every invocation so we cache the data with a TTL
func GetMetadata(remoteUri, cache string, cacheTTL time.Duration, logger *koan.Logger) ([]byte, error) {
	var data []byte

	// check cached exist
	if haveCached(cache, cacheTTL) {
		data, err := os.ReadFile(cache)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(remoteUri)
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
	if err := os.WriteFile(cache, data, 0700); err != nil {
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
