package handlers

import (
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/metadata"
	"strconv"
)

func ListPluginVersions(meta *metadata.RssMetadata, args []string) (string, error) {

	//tagInfo := ""
	output := `Plugin information:
  ID: %d
  Name: %s
  Description: %s
  URL: %s	
  Latest: %s (requires Morpheus %s)

Version History (with min Morpheus):
%s  
`

	var plugin string
	var p metadata.Item
	var semVer, morphVer, pubDate []string

	// check there is id/name argument
	if len(args) < 2 {
		return "", internal.ERR_NO_PLUGIN
	}
	plugin = args[1]

	// get template info using either id or key
	id, err := strconv.Atoi(plugin)
	if err != nil {
		// text, can't convert
		p, id, semVer, morphVer, pubDate, err = meta.GetPluginByName(plugin)
		if err != nil {
			return "", err
		}
	} else {
		// id, converts
		p, semVer, morphVer, pubDate, err = meta.GetPluginByIndex(id)
		if err != nil {
			return "", err
		}
	}

	// create the version info
	verTemplate := "  %s (> %s), published %s\n"
	verOutput := ""

	// we wantt latest first
	for i := len(semVer) - 1; i >= 0; i-- {
		verOutput += fmt.Sprintf(verTemplate, semVer[i], morphVer[i], pubDate[i])
	}

	//for i := range semVer {
	//	verOutput += fmt.Sprintf(verTemplate, semVer[i], morphVer[i], pubDate[i])
	//}
	output = fmt.Sprintf(output, id, p.Code, p.Description, p.FileLink, semVer[len(semVer)-1:][0], morphVer[len(morphVer)-1:][0], verOutput)

	return output, nil
}
