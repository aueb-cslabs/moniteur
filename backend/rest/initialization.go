package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
)

// Map that connects rooms and announcements
var Calendar *types.Calendar

// Initialize Method
func Initialize(initCalendar *types.Calendar) {
	Calendar = initCalendar
}
