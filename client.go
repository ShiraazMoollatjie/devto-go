package devtogo

import "net/http"

const ()

// Client makes all the API calls to dev.to.
type Client struct {
	baseURL string
	http    *http.Client
}

// Option allows the client to be configured with different options.
type Option func(*Client)

func withBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// NewClient creates a dev.to client with the provided options.
func NewClient(opts ...Option) *Client {
	res := &Client{
		baseURL: "http://dev.to/api",
		http:    http.DefaultClient,
	}
	for _, o := range opts {
		o(res)
	}

	return res
}

func getRequest(method, url string, json interface{}) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}
