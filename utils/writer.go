package utils

import (
	"github.com/aueb-cslabs/moniteur/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// UpdateUsers updates the authorized users in config.yml
func UpdateUsers(users map[string]string) error {
	config, _ := types.LoadConfiguration("config/config.yml")
	config.AuthorizedUsers = users
	return write(config, "config/config.yml")
}

func UpdateCalendar(calendar *types.Calendar) error {
	return write(calendar, "config/calendar.yml")
}

// WriteAnnouncement writes announcement to existing.yml
func WriteAnnouncement(msg *types.Announcement) error {
	existing, _ := Read()
	existing.Announcement = msg
	return write(existing, "config/existing.yml")
}

// WriteComment writes comment to existing.yml
func WriteComment(msg *types.Announcement) error {
	existing, _ := Read()
	existing.Comment = msg
	return write(existing, "config/existing.yml")
}

// WriteRoomAnnouncement writes room announcements to existing.yml
func WriteRoomAnnouncement(roomAnnouncements map[string]*types.Announcement) error {
	existing, _ := Read()
	if len(existing.RoomAnnouncements) == 0 {
		existing.RoomAnnouncements = make(map[string]*types.Announcement)
	}
	existing.RoomAnnouncements = roomAnnouncements
	return write(existing, "config/existing.yml")
}

// write an interface and marshals it as a yml file
func write(msg interface{}, file string) error {
	byt, err := yaml.Marshal(msg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, byt, 0666)
}
