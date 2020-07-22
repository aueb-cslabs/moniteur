// +build linux darwin

package main

import (
	"errors"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

type Lesson struct {
	Semester        string `json:"semester"`
	LessonComments  string `json:"Lesson_comments"`
	Room            string `json:"Room"`
	LessonTitle     string `json:"Lesson_title"`
	Professor       string `json:"professor"`
	Time            string `json:"time"`
	Day             string `json:"Day"`
	DepartmentTitle string `json:"Department_title"`
}

type RoomMap struct {
	Rooms        map[string]string `yaml:"rooms,omitempty"`
	RoomsInverse map[string]string `yaml:"-"`
}

var (
	// PLUGIN The plugin to be read by the moniteur agent.
	PLUGIN = Plugin{}
	// Provides map for english-greek room name translation
	mapping = &RoomMap{}
	// The xlsx file that AUEB provides that contains the exams schedule
	exams []byte
	// The link that corresponds to the exams file
	link string
)

type Plugin struct {
}

// Initialize Method that initializes crucial functions for the plugin
func (Plugin) Initialize(examsLink string) {
	if len(mapping.Rooms) == 0 {
		mapping, _ = loadMapping("mapping.yml")
	}
	link = examsLink
	exams, _ = download()
	configureLink()
	exams, _ = download()
	go scheduleExamsDownload()

	ldapConf = LdapConfiguration{}
	if err := loadLdapConfig(&ldapConf); err != nil {
		panic(err)
	}

}

// Schedule Method that returns current schedule from Schedule Master
func (Plugin) Schedule() (*types.Schedule, error) {
	return retriever(), nil
}

// ScheduleRoom Method that returns current schedule and the room that corresponds to it
func (Plugin) ScheduleRoom(room string) (*types.Schedule, error, string) {
	room, changed := checkMapping(room)
	if !changed {
		room = convertChars(room)
	}
	return retriever(), nil, room
}

// ExamsSchedule Method that returns current exams schedule if that exists
func (Plugin) ExamsSchedule() (*types.Schedule, error) {
	schedule := getExamsSchedule()
	var err error
	if len(exams) == 0 {
		err = errors.New("exams schedule not available")
	}
	return schedule, err
}

func (Plugin) ConvertNameInverse(name string) string {
	name, changed := checkMappingInverse(name)
	if !changed {
		name = convertCharsInverse(name)
	}
	return name
}

// ExamsScheduleRoom Method that returns current exams schedule if that exists and the room that corresponds to it
func (Plugin) ExamsScheduleRoom(room string) (*types.Schedule, error, string) {
	room, changed := checkMapping(room)
	if !changed {
		room = convertChars(room)
	}
	schedule := getExamsSchedule()
	var err error
	if len(exams) == 0 {
		err = errors.New("exams schedule not available")
	}
	return schedule, err, room
}

func (Plugin) AuthorizeUser(username string, password string) (bool, error) {
	return authenticateLdap(username, password)
}

func (Plugin) GetRooms() []string {
	rooms := make(map[string]int)
	slots := getEntireSchedule()
	for _, lesson := range slots {
		if rooms[lesson.Room] != 1 {
			rooms[lesson.Room] = 1
		}
	}
	keys := make([]string, 0, len(rooms))
	for key := range rooms {
		keys = append(keys, key)
	}
	return keys
}
