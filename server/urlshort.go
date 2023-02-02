package server

import (
	"net/http"
)

type Store interface {
	GetByKey(short string) (url string, found bool)
}

type Server struct {
	store    Store
	fallback http.Handler
}

func New(store Store, fallback http.Handler) Server {
	return Server{store, fallback}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	reqUrl := request.URL.Path
	val, exists := s.store.GetByKey(reqUrl)
	if exists {
		writer.Header().Set("Location", val)
		writer.WriteHeader(http.StatusMovedPermanently)
	} else {
		s.fallback.ServeHTTP(writer, request)
	}

}
