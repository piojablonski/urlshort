package server_test

import (
	"github.com/piojablonski/urlshort/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var redirects = map[string]string{
	"/tests-for-http": "https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#http.listenandserve-5000-...",
}

type SpyHandler struct {
	calledServeHttp bool
}

func (s *SpyHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	s.calledServeHttp = true
}

func createReqRes(url string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res := httptest.NewRecorder()
	return req, res
}

func TestMapHandler(t *testing.T) {
	t.Run("when called with known url, it the handler redirect to it", func(t *testing.T) {
		srv := server.New(server.NewInmemoryStore(redirects), &SpyHandler{})
		url := "/tests-for-http"
		req, res := createReqRes(url)
		srv.ServeHTTP(res, req)
		assert.Equal(t, res.Code, http.StatusMovedPermanently, "should return status \"moved permanently\"")
		assert.Equal(t, res.Header().Get("Location"), redirects[url], "should have \"Location\" header")
	})

	t.Run("uses fallback handler when path doesn't match any url in a map", func(t *testing.T) {
		fallback := &SpyHandler{}
		srv := server.New(server.NewInmemoryStore(redirects), fallback)
		req, res := createReqRes("/not-existing-path")
		srv.ServeHTTP(res, req)
		assert.True(t, fallback.calledServeHttp, "didn't call ServerHTTP on SpyHandler")
	})
}

const yaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution

`

func TestYamlHandler(t *testing.T) {
	req, res := createReqRes("/urlshort")
	yamlStore := server.NewYamlStore(strings.NewReader(yaml))
	srv := server.New(yamlStore, &SpyHandler{})
	srv.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusMovedPermanently)
	assert.Equal(t, res.Header().Get("Location"), "https://github.com/gophercises/urlshort")
}
