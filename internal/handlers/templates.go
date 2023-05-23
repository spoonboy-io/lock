package handlers

import (
	"fmt"
	"github.com/spoonboy-io/lock/internal"
	"github.com/spoonboy-io/lock/internal/metadata"
	"strings"
)

// ListTemplates will process metadata and provide output
// related to the available starter repositories
func ListTemplates(meta *metadata.Metadata, args []string) (string, error) {
	// handle optional flags --category and --morpheus
	var filterCat, filterMorph string

	var rowString = "%s  %s  %s  %s  %s  %s\n"

	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--category=") {
			filterCat = strings.TrimPrefix(v, "--category=")
		}
		if strings.HasPrefix(v, "--morpheus=") {
			filterMorph = strings.TrimPrefix(v, "--morpheus=")
		}
	}

	// we need these for consistent formatting
	var maxId, maxName, maxCat, maxDesc, maxVer, maxTags int
	for i, p := range *meta {

		// check filters
		if filterCat != "" && p.Category != filterCat {
			continue
		}
		if filterMorph != "" && p.MinimumMorpheus != filterMorph {
			continue
		}

		id := fmt.Sprintf("%d.", i+1)
		// keep track of lengths for header
		if len(id) > maxId {
			maxId = len(id)
		}

		// TODO strictly this doesn't deal with UTF 8
		if len(p.Name) > maxName {
			maxName = len(p.Name)
		}
		if len(p.Category) > maxCat {
			maxCat = len(p.Category)
		}

		lenDesc := len(p.Description)
		if len(p.Description) > 40 {
			lenDesc = 42
		}
		if lenDesc > maxDesc {
			maxDesc = lenDesc
		}

		if len(p.MinimumMorpheus) > maxVer {
			maxVer = len(p.MinimumMorpheus)
		}
		if len(p.Tags) > maxTags {
			maxTags = len(p.Tags)
		}
	}

	// now we know our string lengths, we can assemble the output
	var output string
	var rowCount int

	// aliases
	var padder = internal.WriteCell
	var title = internal.WriteTitle
	var line = internal.Writeline

	for i, p := range *meta {
		// check filters
		if filterCat != "" && p.Category != filterCat {
			continue
		}
		if filterMorph != "" && p.MinimumMorpheus != filterMorph {
			continue
		}

		// lets set a max length on description
		if len(p.Description) > 40 {
			p.Description = cutString(p.Description, 40)
		}
		id := fmt.Sprintf("%d.", i+1)
		output += fmt.Sprintf(rowString, padder(id, maxId), padder(p.Name, maxName), padder(p.Category, maxCat), padder(p.Description, maxDesc), padder(p.MinimumMorpheus, maxVer), padder(p.Tags, maxTags))
		rowCount++
	}

	// add header
	idH := title("ID", maxId)
	idU := line(maxId)
	nameH := title("NAME", maxName)
	nameU := line(maxName)
	catH := title("CAT", maxCat)
	catU := line(maxCat)
	descH := title("DESCRIPTION", maxDesc)
	descU := line(maxDesc)
	minH := title("MIN", maxVer)
	minU := line(maxVer)
	tagH := title("TAGS", maxTags)
	tagU := line(maxTags)

	header1 := ""
	header2 := ""
	if output == "" {
		output = "No templates found.\n"
	} else {
		header1 = fmt.Sprintf(rowString, idH, nameH, catH, descH, minH, tagH)
		header2 = fmt.Sprintf(rowString, idU, nameU, catU, descU, minU, tagU)
	}

	output = fmt.Sprintf("%s%s%s\n", header1, header2, output)
	return output, nil
}

func cutString(data string, cutAt int) string {
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
