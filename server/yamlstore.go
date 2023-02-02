package server

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
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

type YamlStore struct {
	redirects map[string]string
}

func NewYamlStore(r io.Reader) Store {
	redirects := parseYaml(r)
	return &YamlStore{redirects}
}

func (s *YamlStore) GetByKey(short string) (url string, found bool) {
	url, found = s.redirects[short]
	return
}
