package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/go-redis/redis/v7"
)

// Map that connects rooms and announcements
var announcements map[string]*types.Announcement
var authorized map[string]*types.AuthTokenClaim
var calendar *types.Calendar
var authorizedUsers map[string]string
var secret string
var redisClient redis.Client

// Initialize Method
func Initialize(sec string, existing *types.Reader, initCalendar *types.Calendar, authorizedUsersInit map[string]string, redis redis.Client) {
	announcements = make(map[string]*types.Announcement)
	authorized = make(map[string]*types.AuthTokenClaim)
	calendar = initCalendar
	authorizedUsers = authorizedUsersInit
	secret = sec
	redisClient = redis
	if existing != nil {
		com = existing.Comment
		message = existing.Announcement
	}
	if len(existing.RoomAnnouncements) != 0 {
		for key, value := range existing.RoomAnnouncements {
			announcements[key] = value
		}
	}
	go checkAnnouncementExpiration()
	go checkCommentExpiration()
	go checkRoomAnnouncementsExpiration()
}
