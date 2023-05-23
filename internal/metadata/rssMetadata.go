package metadata

import (
	"encoding/xml"
	"errors"
	"github.com/spoonboy-io/koan"
)

// RssMetadata is a representation of the parsed XML RSS metadata
type RssMetadata struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel captures RSS channel info it is not of much interest to us but included for completeness
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	PubDate     string `xml:"pubDate"`
	Items       []Item `xml:"item"`
}

// Item represents a single <item></item> element from the feed, again we capture it all for completeness
// TODO are there more properties which could be included in he feed, such as repository, provider type, and tags
type Item struct {
	Name                string `xml:"name"`
	Code                string `xml:"code"`
	WebLink             string `xml:"webLink"`
	FileLink            string `xml:"fileLink"`
	FileName            string `xml:"fileName"`
	Version             string `xml:"version"`
	MinApplianceVersion string `xml:"minApplianceVersion"`
	MaxApplianceVersion string `xml:"maxApplianceVersion"`
	Description         string `xml:"description"`
	PubDate             string `xml:"pubDate"`
}

var ERR_UNMARSHALLING_RSS = errors.New("could not unmarshal the RSS XML")

// ParseMetadataXML parses the XML RSS feed to the Metadata type
func ParseMetadataXML(data []byte, logger *koan.Logger) (RssMetadata, error) {
	var rssMeta RssMetadata
	err := xml.Unmarshal(data, &rssMeta)
	if err != nil {
		logger.Error("unmarshalling XML RSS", err)
		return rssMeta, ERR_UNMARSHALLING_RSS
	}

	return rssMeta, nil
}
