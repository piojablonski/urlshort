package server

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return nil
}

func DefaultMapHandler(pathsToUrls map[string]string) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		reqUrl := request.URL.Path
		val := pathsToUrls[reqUrl]
		writer.WriteHeader(http.StatusMovedPermanently)

		writer.Header().Set("Location", val)
	}
}
