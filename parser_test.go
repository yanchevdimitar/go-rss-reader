package reader

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yanchevdimitar/go-rss-reader/dto"
)

func TestParser(t *testing.T) {
	tests := []struct {
		Name   string
		XML    string
		Expect func() (items []dto.RssItem, err error)
	}{
		{
			Name: "Ok",
			XML: `<rss version="2.0">
							<channel>
								<item>
									<title><![CDATA[Test]]></title>
									<description><![CDATA[Test]]></description>
									<link>https://Test</link>
		                           <source url="http://Test/">Test</source>
									<pubDate>Sat, 16 Apr 2022 14:43:29 GMT</pubDate>
								</item>
								</channel>
						</rss>`,
			Expect: func() (items []dto.RssItem, err error) {
				return []dto.RssItem{
					{
						Title:       "Test",
						Link:        "https://Test",
						Source:      "Test",
						SourceURL:   "http://Test/",
						Description: "Test",
					},
				}, nil
			},
		},
		{
			Name: "Wrong XML",
			XML:  `<rss ve`,
			Expect: func() (items []dto.RssItem, err error) {
				return nil, ErrXML
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			r := ioutil.NopCloser(bytes.NewReader([]byte(tc.XML)))
			res := &http.Response{
				StatusCode: 200,
				Body:       r,
			}
			result, actualErr := NewDefaultParser().Read(res)
			expected, expectErr := tc.Expect()
			assert.Equal(t, expected, result)
			assert.Equal(t, expectErr, actualErr)
		})
	}
}
