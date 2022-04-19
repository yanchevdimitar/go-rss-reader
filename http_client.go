package reader

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var httpClient HTTPClient

func setHTTPClient(client HTTPClient) {
	httpClient = client
}

func getHTTPClient() HTTPClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return httpClient
}
