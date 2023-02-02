package main

import (
	"fmt"
	"github.com/piojablonski/urlshort/server"
	"log"
	"net/http"
)

func showMessageNotFound(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "the short path not found")
}

func main() {
	srv := server.New(server.NewRedisStore(), http.HandlerFunc(showMessageNotFound))
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
	}
}
