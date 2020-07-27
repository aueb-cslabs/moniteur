package test

import (
	"github.com/aueb-cslabs/moniteur/backend/databases"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

func TestFactory() (e *types.Context) {
	psql := databases.InitializePostgres()
	redisClient, authUsers, tokens := databases.InitializeRedis(types.Redis{
		Ann_DB:    0,
		Users_DB:  1,
		Tokens_DB: 2,
	})

	ctx := &types.Context{
		Context:     nil,
		Calendar:    nil,
		DB:          psql,
		RedisClient: redisClient,
		AuthUsers:   authUsers,
		Tokens:      tokens,
		Secret:      "secret",
	}

	return ctx
}
