package devtogo

import (
	"fmt"
	"time"
)

// PublishedArticle returns a published article with post content for the provided article id.
// https://docs.dev.to/api/#operation/getArticleById
func (c *Client) PublishedArticle(id int32) (*Article, error) {
	var res Article
	err := c.get(c.baseURL+fmt.Sprintf("/articles/%d", id), &res)

	return &res, err
}

// Articles returns a slice of articles according to https://docs.dev.to/api/#operation/getArticles.
func (c *Client) Articles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles?"+qp, &res)

	return res, err
}

// GetMyArticles returns a slice of articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1me/get.
func (c *Client) GetMyArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me?"+qp, &res)

	return res, err
}

// GetMyPublishedArticles returns a slice of published articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1me~1published/get.
func (c *Client) GetMyPublishedArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/published?"+qp, &res)

	return res, err
}

// GetMyUnpublishedArticles returns a slice of unpublished articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1me~1unpublished/get.
func (c *Client) GetMyUnpublishedArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/unpublished?"+qp, &res)

	return res, err
}

// GetAllMyArticles returns a slice of unpublished articles according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1me~1all/get.
func (c *Client) GetAllMyArticles(args Arguments) (Articles, error) {
	var res Articles
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/articles/me/all?"+qp, &res)

	return res, err
}

// CreateArticle creates a post on dev.to according to https://docs.dev.to/api/#tag/articles/paths/~1articles/post.
func (c *Client) CreateArticle(req CreateArticle) (Article, error) {
	var res Article
	err := c.post(c.baseURL+"/articles", ArticleReq{Article: req}, &res)

	return res, err
}

// Update creates a put on dev.to according to https://docs.dev.to/api/#tag/articles/paths/~1articles~1{id}/put
func (c *Client) UpdateArticle(id int, req CreateArticle) (Article, error) {
	var res Article
	err := c.put(c.baseURL+fmt.Sprintf("/articles/%d", id), ArticleReq{Article: req}, &res)

	return res, err
}

// The structs in this file was generated via https://mholt.github.io/json-to-go/.

// ArticleReq is a container type to create articles.
type ArticleReq struct {
	Article CreateArticle `json:"article"`
}

// CreateArticle is a request struct that creates an article.
type CreateArticle struct {
	Title        string   `json:"title"`
	Published    bool     `json:"published"`
	BodyMarkdown string   `json:"body_markdown"`
	Tags         []string `json:"tags"`
	Series       string   `json:"series,omitempty"`
	CanonicalURL string   `json:"canonical_url"`
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
	PublishedTimestamp     emptyTime    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization"`
	Published              bool         `json:"published"`
	Markdown               string       `json:"body_markdown"`
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
	PublishedAt            emptyTime   `json:"published_at"`
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
