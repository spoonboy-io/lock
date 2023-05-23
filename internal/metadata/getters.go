package metadata

import (
	"errors"
)

var (
	ERR_TEMPLATE_ID_NOT_FOUND   = errors.New("no template with that id found")
	ERR_TEMPLATE_NAME_NOT_FOUND = errors.New("no template with that name found")

	ERR_PLUGIN_ID_NOT_FOUND   = errors.New("no plugin with that id found")
	ERR_PLUGIN_NAME_NOT_FOUND = errors.New("no plugin with that name found")
)

// GetTemplateByName will iterate the metadata to retrieve by name key we also return the index as useful
func (md *Metadata) GetTemplateByName(key string) (Plugin, int, error) {
	id := 0
	for _, p := range *md {
		id++
		if p.Name == key {
			return p.Plugin, id, nil
		}
	}
	return Plugin{}, id, ERR_TEMPLATE_NAME_NOT_FOUND
}

// GetTemplateByIndex will iterate the metadata to retrieve by index
func (md *Metadata) GetTemplateByIndex(id int) (Plugin, error) {
	for i, p := range *md {
		if id == i+1 {
			return p.Plugin, nil
		}
	}
	return Plugin{}, ERR_TEMPLATE_ID_NOT_FOUND
}

// GetPluginByName will iterate the rss metadata to retrieve by name key we also return the index as useful
// TODO five return values because it evolved, need to refactor
func (md *RssMetadata) GetPluginByName(key string) (Item, int, []string, []string, []string, []string, error) {
	rowCount := 0
	lastCode := ""

	for _, p := range md.Channel.Items {
		if p.Code != lastCode {
			rowCount++
		}

		if p.Code == key {
			semVar, morphVar, pubDate, fileName := getPluginVersions(md, p.Code)
			return p, rowCount, semVar, morphVar, pubDate, fileName, nil
		}

		lastCode = p.Code
	}

	return Item{}, rowCount, []string{}, []string{}, []string{}, []string{}, ERR_PLUGIN_NAME_NOT_FOUND
}

// GetPluginByIndex will iterate the rss metadata to retrieve by index
func (md *RssMetadata) GetPluginByIndex(id int) (Item, []string, []string, []string, []string, error) {
	rowCount := 0
	lastCode := ""

	for _, p := range md.Channel.Items {
		if p.Code != lastCode {
			rowCount++
		}

		if id == rowCount {
			semVar, morphVar, pubDate, fileName := getPluginVersions(md, p.Code)
			return p, semVar, morphVar, pubDate, fileName, nil
		}

		lastCode = p.Code
	}

	return Item{}, []string{}, []string{}, []string{}, []string{}, ERR_PLUGIN_ID_NOT_FOUND
}

// helper to iterate the rss meta again and collect versions
// TODO bit messy with four returns, we'll return a struct when we refactor
func getPluginVersions(meta *RssMetadata, code string) ([]string, []string, []string, []string) {
	var semVer, morphVer, pubDate, fileName []string
	for _, p := range meta.Channel.Items {
		if p.Code == code {
			semVer = append(semVer, p.Version)
			morphVer = append(morphVer, p.MinApplianceVersion)
			pubDate = append(pubDate, p.PubDate)
			fileName = append(fileName, p.FileName)
		}
	}
	return semVer, morphVer, pubDate, fileName
}
