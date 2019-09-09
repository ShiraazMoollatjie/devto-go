package devtogo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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

func getRequest(method, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func (c *Client) Get(url string, target interface{}) error {
	req, err := getRequest(http.MethodGet, url)
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
