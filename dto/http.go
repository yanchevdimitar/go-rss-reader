package dto

import (
	"net/http"
)

type URLResponse struct {
	URL      string
	Response *http.Response
	Error    error
}

type RssResponse struct {
	Items map[string][]RssItem
}
