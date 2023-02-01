package server

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
)

func parseYaml(r io.Reader) map[string]string {
	var data []struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}
	err := yaml.NewDecoder(r).Decode(&data)
	if err != nil {
		log.Print(err)
	}

	var res = make(map[string]string)

	for _, v := range data {
		res[v.Path] = v.Url
	}
	return res

}

func YamlHandler(r io.Reader, fallback http.Handler) http.HandlerFunc {
	redirects := parseYaml(r)
	return MapHandler(redirects, fallback)
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		reqUrl := request.URL.Path
		val, exists := pathsToUrls[reqUrl]
		if exists {
			writer.Header().Set("Location", val)
			writer.WriteHeader(http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}
