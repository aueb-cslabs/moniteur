package databases

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/go-redis/redis/v7"
	"log"
)

func InitializeRedis(config types.Redis) (*redis.Client, *redis.Client) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Ann_DB,
	})
	authUsers := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Users_DB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	_, err = authUsers.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	return redisClient, authUsers
}
