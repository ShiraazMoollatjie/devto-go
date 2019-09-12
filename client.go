package devtogo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Client makes all the API calls to dev.to.
type Client struct {
	apiKey  string
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

// WithApiKey sets the dev.to api key to use for this client.
// see https://docs.dev.to/api/#section/Authentication for how to set one up.
func WithApiKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
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

func (c *Client) getRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Add("api-key", c.apiKey)
	}

	return req, err
}

// Get returns an error if the http client cannot perform a HTTP GET for the provided URL.
func (c *Client) Get(url string, target interface{}) error {
	req, err := c.getRequest(http.MethodGet, url)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error from dev.to api")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &target)
}
