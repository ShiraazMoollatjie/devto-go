package devtogo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	NewClient(withBaseURL(ts.URL))
	fmt.Println(ts.URL)
	assert.True(t, strings.Contains(ts.URL, "127.0.0.1"))
}

func TestNewClientApiKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "myApiKey", r.Header.Get("api-key"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	fmt.Println(ts.URL)
	assert.True(t, strings.Contains(ts.URL, "127.0.0.1"))
}
