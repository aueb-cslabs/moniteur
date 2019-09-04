package types

// Announcement contains the message to broadcast and the date that it expires
type Announcement struct {
	End string `json:"end"`
	Msg string `json:"msg"`
}
