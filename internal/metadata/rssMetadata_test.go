package metadata_test

import (
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/lock/internal/metadata"
	"testing"
)

func TestParseMetadataXML(t *testing.T) {
	// single case just to make sure we have the
	// correct data

	logger := &koan.Logger{}
	testCases := []struct {
		name            string
		xml             []byte
		wantName        string
		wantCode        string
		wantWebLink     string
		wantFileLink    string
		wantFileName    string
		wantVersion     string
		wantMinVersion  string
		wantMaxVersion  string
		wantDescription string
		wantPubDate     string
		wantErr         error
	}{
		{"simple good format",
			[]byte(`
<rss xmlns:content="http://purl.org/rss/1.0/modules/content/" version="2.0">
<channel>
<item>
<name>Efficient IP</name>
<code>efficient-ip-plugin</code>
<webLink>https://share.morpheusdata.com/api/packages/pull?code=efficient-ip-plugin</webLink>
<fileLink>https://share.morpheusdata.com/efficient-ip-plugin</fileLink>
<fileName>morpheus-efficientip-plugin-1.0.0-all.jar</fileName>
<version>1.0.0</version>
<minApplianceVersion>5.4.5</minApplianceVersion>
<maxApplianceVersion/>
<description>Plugin for Efficient IP</description>
<pubDate>Thu, 28 Apr 2022 15:11:24 GMT</pubDate>
</item>
	</channel>
	</rss>
`),

			"Efficient IP",
			"efficient-ip-plugin",
			"https://share.morpheusdata.com/api/packages/pull?code=efficient-ip-plugin",
			"https://share.morpheusdata.com/efficient-ip-plugin",
			"morpheus-efficientip-plugin-1.0.0-all.jar",
			"1.0.0",
			"5.4.5",
			"",
			"Plugin for Efficient IP",
			"Thu, 28 Apr 2022 15:11:24 GMT",
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := metadata.ParseMetadataXML(tc.xml, logger)
			if gotErr != tc.wantErr {
				t.Errorf("wanted %v got %v", tc.wantErr, gotErr)
			}

			if gotRes.Channel.Items[0].Name != tc.wantName {
				t.Errorf("wanted:%v, got %v", tc.wantName, gotRes.Channel.Items[0].Name)
			}

			if gotRes.Channel.Items[0].Code != tc.wantCode {
				t.Errorf("wanted:%v, got %v", tc.wantCode, gotRes.Channel.Items[0].Code)
			}

			if gotRes.Channel.Items[0].WebLink != tc.wantWebLink {
				t.Errorf("wanted:%v, got %v", tc.wantWebLink, gotRes.Channel.Items[0].WebLink)
			}

			if gotRes.Channel.Items[0].FileLink != tc.wantFileLink {
				t.Errorf("wanted:%v, got %v", tc.wantFileLink, gotRes.Channel.Items[0].FileLink)
			}

			if gotRes.Channel.Items[0].FileName != tc.wantFileName {
				t.Errorf("wanted:%v, got %v", tc.wantFileName, gotRes.Channel.Items[0].FileName)
			}

			if gotRes.Channel.Items[0].Version != tc.wantVersion {
				t.Errorf("wanted:%v, got %v", tc.wantVersion, gotRes.Channel.Items[0].Version)
			}

			if gotRes.Channel.Items[0].MinApplianceVersion != tc.wantMinVersion {
				t.Errorf("wanted:%v, got %v", tc.wantMinVersion, gotRes.Channel.Items[0].MinApplianceVersion)
			}

			if gotRes.Channel.Items[0].MaxApplianceVersion != tc.wantMaxVersion {
				t.Errorf("wanted:%v, got %v", tc.wantMaxVersion, gotRes.Channel.Items[0].MaxApplianceVersion)
			}

			if gotRes.Channel.Items[0].Description != tc.wantDescription {
				t.Errorf("wanted:%v, got %v", tc.wantDescription, gotRes.Channel.Items[0].Description)
			}

			if gotRes.Channel.Items[0].PubDate != tc.wantPubDate {
				t.Errorf("wanted:%v, got %v", tc.wantPubDate, gotRes.Channel.Items[0].PubDate)
			}
		})
	}
}
