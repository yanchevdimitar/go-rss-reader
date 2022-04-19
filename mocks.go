package reader

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/yanchevdimitar/go-rss-reader/dto"
)

type MockRssParser struct {
	mock.Mock
}

func (m *MockRssParser) Read(resp *http.Response) (items []dto.RssItem, err error) {
	args := m.Called(resp)
	return args.Get(0).([]dto.RssItem), args.Error(1)
}

type MockRssReader struct {
	mock.Mock
}

func (m *MockRssReader) Parse() (rss []dto.RssResponse, err error) {
	args := m.Called()
	return args.Get(0).([]dto.RssResponse), args.Error(1)
}

type HTTPClientMock struct {
	mock.Mock
}

func (hc *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	args := hc.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
