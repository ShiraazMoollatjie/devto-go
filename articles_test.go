package devtogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "article.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/article/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.GetArticle(167919)
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestGetArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	tests := []struct {
		name                string
		arguments           Arguments
		expectedQueryParams string
	}{
		{"No params", Defaults(), ""},
		{"Page param", Arguments{"page": "1"}, "page=1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/articles?"+test.expectedQueryParams, r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			client := NewClient(withBaseURL(ts.URL))
			articles, err := client.GetArticles(test.arguments)
			assert.NoError(t, err)
			assert.Equal(t, res, articles)
		})
	}
}

func TestGetMyArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetMyArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetMyPublishedArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/published?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetMyPublishedArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetMyUnpublishedArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/unpublished?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetMyUnpublishedArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestGetAllMyArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "myarticles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/me/all?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetAllMyArticles(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestCreateArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticle{
		Tags:         []string{"go", "help"},
		Series:       "api",
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.CreateArticle(testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestCreateArticleNoSeriesField(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticle{
		Tags:         []string{"go", "help"},
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.CreateArticle(testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func TestUpdateArticle(t *testing.T) {
	var res Article
	b := unmarshalGoldenFileBytes(t, "create_article.json", &res)
	testArticle := CreateArticle{
		Tags:         []string{"go", "help"},
		Series:       "api",
		Published:    false,
		BodyMarkdown: "This is some markdown",
		Title:        "My First Post via API",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles/1000", r.URL.Path)
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)

		var car ArticleReq
		assert.NoError(t, json.Unmarshal(rb, &car))
		assert.Equal(t, ArticleReq{Article: testArticle}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.UpdateArticle(1000, testArticle)
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func unmarshalGoldenFileBytes(t *testing.T, filename string, payload interface{}) []byte {
	p := filepath.Join("testdata", filename)
	b, err := ioutil.ReadFile(p)
	assert.NoError(t, err)

	err = json.Unmarshal(b, &payload)
	assert.NoError(t, err)

	return b
}
