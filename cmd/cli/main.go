package main

import (
	"fmt"
	"github.com/piojablonski/urlshort/server"
	"log"
	"net/http"
)

func doNothing(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "the short path not found")
}

func main() {

	var redirects = map[string]string{
		"/tests-for-http": "https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#http.listenandserve-5000-...",
	}
	err := http.ListenAndServe(":8080",
		server.MapHandler(
			redirects,
			http.HandlerFunc(doNothing),
		))
	if err != nil {
		log.Fatal(err)
	}
}
