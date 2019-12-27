package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/go-redis/redis/v7"
)

// Map that connects rooms and announcements
var authorized map[string]*types.AuthTokenClaim
var calendar *types.Calendar
var authorizedUsers map[string]string
var secret string
var redisClient redis.Client

// Initialize Method
func Initialize(sec string, initCalendar *types.Calendar, authorizedUsersInit map[string]string, redis redis.Client) {
	authorized = make(map[string]*types.AuthTokenClaim)
	calendar = initCalendar
	authorizedUsers = authorizedUsersInit
	secret = sec
	redisClient = redis
}
