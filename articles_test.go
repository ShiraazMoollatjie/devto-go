package devtogo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestGetArticle(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "article.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/article/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.GetArticle("167919")
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestGetArticles(t *testing.T) {
	var res Articles
	b := unmarshalGoldenFileBytes(t, "articles.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetArticles()
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}

func unmarshalGoldenFileBytes(t *testing.T, filename string, payload interface{}) []byte {
	p := filepath.Join("testdata", filename)
	b, err := ioutil.ReadFile(p)
	err = json.Unmarshal(b, &payload)
	assert.NoError(t, err)

	return b
}
