package rest

import "github.com/aueb-cslabs/moniteur/types"

// Map that connects rooms and announcements
var announcements map[string]*types.Announcement
var authorized map[string]*types.AuthTokenClaim
var secret string

// Initialize Method
func Initialize(sec string) {
	announcements = make(map[string]*types.Announcement)
	authorized = make(map[string]*types.AuthTokenClaim)
	secret = sec
	go checkAnnouncementExpiration()
	go checkCommentExpiration()
	go checkRoomAnnouncementsExpiration()
}
