package handlers

import (
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/metadata"
	"strings"
)

// ListPlugins will process metadata and provide output
// related to the available starter plugin jar versions
func ListPlugins(meta *metadata.RssMetadata, args []string) (string, error) {
	// handle optional flag --morpheus
	var filterMorph string

	var rowString = "%s  %s  %s  %s  %s\n"

	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--morpheus=") {
			filterMorph = strings.TrimPrefix(v, "--morpheus=")
		}
	}

	_ = filterMorph

	var output string
	var rowCount int
	var lastCode string
	var semVer, morphVer []string

	// we make a first pass for widths
	var maxId, maxCode, maxDesc, maxLatestVer, maxLatestMorph int
	for _, p := range meta.Channel.Items {
		id := fmt.Sprintf("%d.", rowCount)
		if p.Code != lastCode {
			if rowCount != 0 {

				if len(id) > maxId {
					maxId = len(id)
				}

				if len(p.Code) > maxCode {
					maxCode = len(p.Code)
				}

				if len(p.Description) > maxDesc {
					maxDesc = len(p.Description)
				}

				if len(semVer[len(semVer)-1:][0]) > maxLatestVer {
					maxLatestVer = len(semVer[len(semVer)-1:][0])
				}

				if len(morphVer[len(morphVer)-1:][0]) > maxLatestMorph {
					maxLatestMorph = len(morphVer[len(morphVer)-1:][0])
				}

				// reset
				semVer = []string{}
				morphVer = []string{}
			}
			rowCount++
		}

		// append to the version slices
		semVer = append(semVer, p.Version)
		morphVer = append(morphVer, p.MinApplianceVersion)
		lastCode = p.Code
	}

	// reset for data pass
	rowCount = 0
	lastCode = ""

	// aliases
	var padder = internal.WriteCell
	var title = internal.WriteTitle
	var line = internal.Writeline

	// data pass
	for _, p := range meta.Channel.Items {
		id := fmt.Sprintf("%d.", rowCount+1)
		if p.Code != lastCode {
			if rowCount != 0 {

			}
			output += fmt.Sprintf(rowString, padder(id, maxId), padder(p.Code, maxCode), padder(p.Description, maxDesc), padder(semVer[len(semVer)-1:][0], maxLatestVer), padder(morphVer[len(morphVer)-1:][0], maxLatestMorph))
			// reset
			semVer = []string{}
			morphVer = []string{}
			rowCount++
		}

		// append to the version slices
		semVer = append(semVer, p.Version)
		morphVer = append(morphVer, p.MinApplianceVersion)
		lastCode = p.Code
	}

	// add header
	idH := title("ID", maxId)
	idU := line(maxId)
	nameH := title("NAME", maxCode)
	nameU := line(maxCode)
	descH := title("DESCRIPTION", maxDesc)
	descU := line(maxDesc)
	minH := title("LATEST", maxLatestVer)
	minU := line(maxLatestVer)
	if maxLatestMorph < 8 {
		maxLatestMorph = 8
	}
	tagH := title("MORPHEUS", maxLatestMorph)
	tagU := line(maxLatestMorph)

	header1 := ""
	header2 := ""
	if output == "" {
		output = "No plugins found.\n"
	} else {
		header1 = fmt.Sprintf(rowString, idH, nameH, descH, minH, tagH)
		header2 = fmt.Sprintf(rowString, idU, nameU, descU, minU, tagU)
	}

	output = fmt.Sprintf("%s%s%s\n", header1, header2, output)
	return output, nil
}
