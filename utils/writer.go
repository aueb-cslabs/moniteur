package utils

import (
	"github.com/aueb-cslabs/moniteur/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// WriteAnnouncement writes announcement to existing.yml
func WriteAnnouncement(msg *types.Announcement) error {
	existing, _ := Read()
	existing.Announcement = msg
	return write(existing)
}

// WriteComment writes comment to existing.yml
func WriteComment(msg *types.Announcement) error {
	existing, _ := Read()
	existing.Comment = msg
	return write(existing)
}

// WriteRoomAnnouncement writes room announcements to existing.yml
func WriteRoomAnnouncement(roomAnnouncements map[string]*types.Announcement) error {
	existing, _ := Read()
	existing.RoomAnnouncements = roomAnnouncements
	return write(existing)
}

// write an interface and marshals it as a yml file
func write(msg interface{}) error {
	byt, err := yaml.Marshal(msg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("existing.yml", byt, 0644)
}
