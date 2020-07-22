package databases

import (
	"log"
	"os"

	"github.com/go-redis/redis/v7"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

func InitializeRedis(config types.Redis) (*redis.Client, *redis.Client, *redis.Client) {
	address := os.Getenv("REDIS_URI")
	password := os.Getenv("REDIS_PASSWORD")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       config.Ann_DB,
	})
	authUsers := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       config.Users_DB,
	})
	tokens := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       config.Tokens_DB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	_, err = authUsers.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	_, err = tokens.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	return redisClient, authUsers, tokens
}
