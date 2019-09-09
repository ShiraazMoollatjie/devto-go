package devtogo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// GetArticles returns a slice of articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles/get.
func (c *Client) GetArticles() (Articles, error) {
	req, err := getRequest(http.MethodGet, c.baseURL+"/articles", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error from dev.to api")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Articles
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// The structs in this file was generated via https://mholt.github.io/json-to-go/.
// Articles represents an article from the dev.to api.
type Articles []struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                []string     `json:"tag_list"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization"`
}

// User represents a user from the dev.to api.
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`
	ProfileImage90  string `json:"profile_image_90"`
}

// Organization represents an organization from the dev.to api.
type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}
