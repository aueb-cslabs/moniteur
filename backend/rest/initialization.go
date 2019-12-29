package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
)

// Map that connects rooms and announcements
var Authorized map[string]*types.AuthTokenClaim
var Calendar *types.Calendar

// Initialize Method
func Initialize(initCalendar *types.Calendar) {
	Authorized = make(map[string]*types.AuthTokenClaim)
	Calendar = initCalendar
}
