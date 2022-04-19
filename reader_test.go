package reader

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/yanchevdimitar/go-rss-reader/dto"
)

func TestReader(t *testing.T) {
	tests := []struct {
		Name               string
		mockResponseReader *MockRssParser
		mockHTTP           *HTTPClientMock
		Expect             func() (items []dto.RssResponse, err error)
	}{
		{
			Name: "Ok",
			mockResponseReader: func() *MockRssParser {
				mrr := &MockRssParser{}
				mrr.On("Read", mock.AnythingOfType("*http.Response")).Return([]dto.RssItem{
					{
						Title:       "Test",
						Link:        "https://Test",
						Source:      "Test",
						SourceURL:   "http://Test/",
						Description: "Test",
					},
				}, nil)

				return mrr
			}(),
			mockHTTP: func() *HTTPClientMock {
				mh := HTTPClientMock{}
				mh.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{Status: "200 OK", StatusCode: 200}, nil)

				return &mh
			}(),
			Expect: func() (items []dto.RssResponse, err error) {
				return []dto.RssResponse{
					{
						map[string][]dto.RssItem{"test": {
							{
								Title:       "Test",
								Link:        "https://Test",
								Source:      "Test",
								SourceURL:   "http://Test/",
								Description: "Test",
							},
						},
						},
					},
				}, nil
			},
		},
		{
			Name: "Error on http request",
			mockResponseReader: func() *MockRssParser {
				mrr := &MockRssParser{}
				mrr.On("Read", mock.AnythingOfType("*http.Response")).Return([]dto.RssItem{
					{
						Title:       "Test",
						Link:        "https://Test",
						Source:      "Test",
						SourceURL:   "http://Test/",
						Description: "Test",
					},
				}, nil)

				return mrr
			}(),
			mockHTTP: func() *HTTPClientMock {
				mh := HTTPClientMock{}
				mh.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{Status: "404", StatusCode: 404}, ErrHTTP)

				return &mh
			}(),
			Expect: func() (items []dto.RssResponse, err error) {
				return nil, ErrHTTP
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			setHTTPClient(tc.mockHTTP)
			result, actualErr := NewDafaultReader([]string{"test"}, tc.mockResponseReader).Parse()
			expected, expectErr := tc.Expect()
			assert.Equal(t, result, expected)
			require.Equal(t, expectErr, actualErr)
		})
	}
}
