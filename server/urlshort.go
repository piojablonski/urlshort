package server

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		reqUrl := request.URL.Path
		val, exists := pathsToUrls[reqUrl]
		if exists {
			writer.WriteHeader(http.StatusMovedPermanently)
			writer.Header().Set("Location", val)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}
