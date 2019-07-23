package main

import (
	"encoding/json"
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
	"io/ioutil"
	"net/http"
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

// PLUGIN The plugin to be read by the moniteur agent.
var PLUGIN = Plugin{}

type Plugin struct {
}

func (Plugin) Schedule() (*types.Schedule, error) {

	resp := &types.Schedule{}

	for _, lesson := range getEntireSchedule() {
		subject := &types.ScheduleSlot{}
		lessonTime := strings.Split(lesson.Time, "-")
		start, _ := strconv.ParseInt(lessonTime[0], 10, 64)
		end, _ := strconv.ParseInt(lessonTime[1], 10, 64)
		subject.Start = start
		subject.End = end
		subject.Room = lesson.Room
		subject.Day = determineDay(lesson.Day)
		subject.Title = lesson.LessonTitle
		subject.Host = lesson.Professor
		resp.Slots = append(resp.Slots, subject)
	}

	return resp, fmt.Errorf("this must be working")
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
