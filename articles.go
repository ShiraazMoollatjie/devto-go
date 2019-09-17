package devtogo

import (
	"fmt"
	"time"
)

// GetArticle returns an article with post content for the provided article id.
// https://docs.dev.to/api/#tag/articles/paths/~1articles~1{id}/get
func (c *Client) GetArticle(id string) (*Article, error) {
	var res Article
	err := c.get(c.baseURL+fmt.Sprintf("/article/%s", id), &res)

	return &res, err
}

// GetArticles returns a slice of articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles/get.
func (c *Client) GetArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles?"+qp, &res)

	return res, err
}

// CreateArticle creates a post on dev.to according to https://docs.dev.to/api/#tag/articles/paths/~1articles/post
func (c *Client) CreateArticle(req CreateArticle) (Article, error) {
	var res Article
	err := c.post(c.baseURL+"/articles", ArticleReq{Article: req}, &res)

	return res, err
}

// Update creates a put on dev.to according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1{id}/put
func (c *Client) UpdateArticle(id string, req CreateArticle) (Article, error) {
	var res Article
	err := c.put(c.baseURL+fmt.Sprintf("/articles/%s", id), ArticleReq{Article: req}, &res)

	return res, err
}

// The structs in this file was generated via https://mholt.github.io/json-to-go/.

// ArticleReq is a container type to create articles.
type ArticleReq struct {
	Article CreateArticle `json:"article"`
}

// CreateArticle is a request struct that creates an article.
type CreateArticle struct {
	Title        string `json:"title"`
	Published    bool   `json:"published"`
	BodyMarkdown string `json:"body_markdown"`
	Tags         string `json:"tags"`
	Series       string `json:"series"`
	CanonicalURL string `json:"canonical_url"`
}

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

// Article represents a single article in the dev.to api. Also has more fields than Articles.
type Article struct {
	TypeOf                 string      `json:"type_of"`
	ID                     int         `json:"id"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	CoverImage             string      `json:"cover_image"`
	ReadablePublishDate    string      `json:"readable_publish_date"`
	SocialImage            string      `json:"social_image"`
	TagList                string      `json:"tag_list"`
	Tags                   []string    `json:"tags"`
	Slug                   string      `json:"slug"`
	Path                   string      `json:"path"`
	URL                    string      `json:"url"`
	CanonicalURL           string      `json:"canonical_url"`
	CommentsCount          int         `json:"comments_count"`
	PositiveReactionsCount int         `json:"positive_reactions_count"`
	CreatedAt              time.Time   `json:"created_at"`
	EditedAt               interface{} `json:"edited_at"`
	CrosspostedAt          interface{} `json:"crossposted_at"`
	PublishedAt            time.Time   `json:"published_at"`
	LastCommentAt          time.Time   `json:"last_comment_at"`
	BodyHTML               string      `json:"body_html"`
	BodyMarkdown           string      `json:"body_markdown"`
	User                   User        `json:"user"`
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
