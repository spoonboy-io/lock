package handlers

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/metadata"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DownloadPluginVersion(meta *metadata.RssMetadata, args []string) (string, error) {
	var plugin, version string
	var p metadata.Item
	var semVer, fileName []string

	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--version=") {
			version = strings.TrimPrefix(v, "--version=")
		} else {
			// anything else should be the template?
			plugin = v
		}
	}

	// if empty we have an problem
	if plugin == "" {
		return "", internal.ERR_NO_PLUGIN
	}

	// get template info using either id or key
	id, err := strconv.Atoi(plugin)
	if err != nil {
		// text, can't convert
		p, id, semVer, _, _, fileName, err = meta.GetPluginByName(plugin)
		if err != nil {
			return "", err
		}
	} else {
		// id, converts
		p, semVer, _, _, fileName, err = meta.GetPluginByIndex(id)
		if err != nil {
			return "", err
		}
	}

	fileTarget := fmt.Sprintf("https://share.morpheusdata.com/api/packages/pull?code=%s", p.Code)
	fName := ""

	// check if requested version
	if version != "" {
		validVer := false

		for i, v := range semVer {
			if version == v {
				validVer = true
				fName = fileName[i]
			}

		}
		if !validVer {
			return "", internal.ERR_INVALID_PLUGIN_VERSION
		}
		fileTarget = fmt.Sprintf("https://share.morpheusdata.com/api/packages/pull?code=%s&version=%s", p.Code, version)
	}

	// deal with no version
	if version == "" {
		version = "latest"
		fileTarget = fmt.Sprintf("https://share.morpheusdata.com/api/packages/pull?code=%s", p.Code)
		fName = fileName[len(fileName)-1:][0]
	}

	folder := filepath.Join(p.Code, version)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return "", err
	}

	resp, err := grab.Get(filepath.Join(folder, fName), fileTarget)
	if err != nil {
		return "", err
	}

	// confirm
	output := `Morpheus plugin downloaded:
  Created at: %s		
  Using plugin: %s
  Using version: %s

`

	output = fmt.Sprintf(output, resp.Filename, p.Code, version)
	return output, nil
}
