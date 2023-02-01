package server_test

import (
	"github.com/piojablonski/urlshort/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var redirects = map[string]string{
	"/tests-for-http": "https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#http.listenandserve-5000-...",
}

func TestMapHandler(t *testing.T) {
	t.Run("when called with known url, it the handler redirect to it", func(t *testing.T) {

		url := "/tests-for-http"
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()

		server.DefaultMapHandler(redirects).ServeHTTP(res, req)
		assert.Equal(t, res.Code, http.StatusMovedPermanently)
		assert.Equal(t, res.Header().Get("Location"), redirects[url])
	})
}
