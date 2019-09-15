package types

// Announcement contains the message to broadcast and the date that it expires
type AnnouncementPost struct {
	End string `json:"end"`
	Msg string `json:"msg"`
}

type Announcement struct {
	End int64  `json:"end" yaml:"end"`
	Msg string `json:"msg" yaml:"msg"`
}
