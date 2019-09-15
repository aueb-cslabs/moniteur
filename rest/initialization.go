package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/aueb-cslabs/moniteur/utils"
)

// Map that connects rooms and announcements
var announcements map[string]*types.Announcement
var authorized map[string]*types.AuthTokenClaim
var secret string

// Initialize Method
func Initialize(sec string, existing *utils.Reader) {
	announcements = make(map[string]*types.Announcement)
	authorized = make(map[string]*types.AuthTokenClaim)
	secret = sec
	com = existing.Comment
	announcements = existing.RoomAnnouncements
	message = existing.Announcement
	go checkAnnouncementExpiration()
	go checkCommentExpiration()
	go checkRoomAnnouncementsExpiration()
}
