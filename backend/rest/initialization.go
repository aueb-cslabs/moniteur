package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
)

// Map that connects rooms and announcements
var Authorized map[string]*types.AuthTokenClaim
var Calendar *types.Calendar
var Secret string

// Initialize Method
func Initialize(sec string, initCalendar *types.Calendar) {
	Authorized = make(map[string]*types.AuthTokenClaim)
	Calendar = initCalendar
	Secret = sec
}
