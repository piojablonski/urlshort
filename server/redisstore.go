package server

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() Store {
	store := &RedisStore{}
	store.client = redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     "localhost:6379",
		Password: "",
	})
	return store
}

func (s *RedisStore) GetByKey(short string) (url string, found bool) {
	url, err := s.client.Get(short).Result()
	if err != nil {
		found = false
		log.Printf("problem fetching data from redis, %v", err)
	} else {
		found = true
	}
	return
}
