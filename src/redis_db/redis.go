package redis_db

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func SetupRedis() {
	RDB = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}