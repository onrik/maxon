package maxon

import (
	"net/http"
)

type Options struct {
	apiURL     string
	httpClient *http.Client
}

type Option func(*Options)

// WithHttpClient - set custom http client
func WithHttpClient(httpClient *http.Client) Option {
	return func(o *Options) {
		o.httpClient = httpClient
	}
}

// WithApiUrl - set api url
func WithApiUrl(apiURL string) Option {
	return func(o *Options) {
		o.apiURL = apiURL
	}
}
