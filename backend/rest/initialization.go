package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/go-redis/redis/v7"
)

// Map that connects rooms and announcements
var authorized map[string]*types.AuthTokenClaim
var calendar *types.Calendar
var secret string
var redisClient, authUsers redis.Client

// Initialize Method
func Initialize(sec string, initCalendar *types.Calendar, authUsers redis.Client, redis redis.Client) {
	authorized = make(map[string]*types.AuthTokenClaim)
	calendar = initCalendar
	secret = sec
	redisClient = redis
}
