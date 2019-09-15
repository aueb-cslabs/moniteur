package types

// Reader struct that contains all announcements from existing.yml
type Reader struct {
	Announcement      *Announcement            `yaml:"announcement"`
	RoomAnnouncements map[string]*Announcement `yaml:"room_announcements"`
	Comment           *Announcement            `yaml:"comment"`
}
