// +build linux darwin

package main

import (
	"encoding/json"
	"github.com/aueb-cslabs/moniteur/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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
	Rooms map[string]string `yaml:"rooms,omitempty"`
}

// PLUGIN The plugin to be read by the moniteur agent.
var PLUGIN = Plugin{}
var mapping = &RoomMap{}

type Plugin struct {
}

func (Plugin) Initialize() {
	if len(mapping.Rooms) == 0 {
		mapping, _ = loadMapping("mapping.yml")
	}
}

func (Plugin) Schedule(room string) (*types.Schedule, error, string) {

	resp := &types.Schedule{}
	for _, lesson := range getEntireSchedule() {
		subject := &types.ScheduleSlot{}
		lessonTime := strings.Split(lesson.Time, "-")
		start, _ := strconv.ParseInt(lessonTime[0], 10, 64)
		end, _ := strconv.ParseInt(lessonTime[1], 10, 64)
		subject.Start = start * int64(3600)
		subject.End = end * int64(3600)
		subject.Room = lesson.Room
		subject.Day = determineDay(lesson.Day)
		subject.Title = lesson.LessonTitle
		subject.Host = lesson.Professor
		resp.Slots = append(resp.Slots, subject)
	}

	room, changed := checkMapping(room)

	if !changed {
		room = convertChars(room)
	}

	return resp, nil, room
}

func convertChars(room string) string {
	re := regexp.MustCompile("[0-9]+")

	if re.MatchString(room) {
		if strings.Contains(room, "a") {
			room = strings.ReplaceAll(room, "a", "Α")
		}
		if strings.Contains(room, "d") {
			room = strings.ReplaceAll(room, "d", "Δ")
		}
		if strings.Contains(room, "t") {
			room = strings.ReplaceAll(room, "t", "Τ")
		}
		return room
	}

	return room
}

func getEntireSchedule() []*Lesson {
	resp, _ := http.Get("http://schedule.aueb.gr/mobile/")
	bts, _ := ioutil.ReadAll(resp.Body)
	var slots []*Lesson
	_ = json.Unmarshal(bts, &slots)
	return slots
}

func determineDay(day string) int {
	switch day {
	case "Δευτέρα":
		return 1
	case "Τρίτη":
		return 2
	case "Τετάρτη":
		return 3
	case "Πέμπτη":
		return 4
	case "Παρασκευή":
		return 5
	case "Σάββατο":
		return 6
	}
	return 0
}

func loadMapping(file string) (*RoomMap, error) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	rooms := &RoomMap{}
	if err := yaml.Unmarshal(byt, rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func checkMapping(room string) (string, bool) {
	rooms := mapping.Rooms

	if rooms != nil && rooms[room] != "" {
		return rooms[room], true
	}

	return room, false
}
