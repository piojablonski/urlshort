package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var redirects = map[string]string{
	"/tests-for-http": "https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#http.listenandserve-5000-...",
	"/g":              "https:://google.com",
	"/gotests":        "https://quii.gitbook.io",
}

func main() {
	fmt.Println("redis tutorial")

	client := redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     "localhost:6379",
		Password: "",
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	for k, v := range redirects {
		status := client.Set(k, v, 0)
		fmt.Println(status)
	}
}
