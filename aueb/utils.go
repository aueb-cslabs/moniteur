package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// loadMapping Method that loads the room mapping from english to greek names
func loadMapping(file string) (*RoomMap, error) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	rooms := &RoomMap{}
	if err := yaml.Unmarshal(byt, rooms); err != nil {
		return nil, err
	}
	rooms.RoomsInverse = make(map[string]string)
	for key, value := range rooms.Rooms {
		rooms.RoomsInverse[value] = key
	}
	return rooms, nil
}

// checkMapping Method that checks the room map in order to retrieve the right room name
func checkMapping(room string) (string, bool) {
	rooms := mapping.Rooms

	if rooms != nil && rooms[room] != "" {
		return rooms[room], true
	}

	return room, false
}

// checkMappingInverse Method that checks the room map in order to retrieve the right room name
func checkMappingInverse(room string) (string, bool) {
	rooms := mapping.RoomsInverse

	if rooms != nil && rooms[room] != "" {
		return rooms[room], true
	}

	return room, false
}

// convertTime Method that converts time to int64
func convertTime(timestamp string) int64 {
	convert := strings.Split(timestamp, ":")
	var hourF float64

	hourInt, _ := strconv.Atoi(convert[0])

	hourF = float64(hourInt)
	if convert[1] == "30" {
		hourF += .5
	}

	return int64(hourF * 3600)
}

// download Method that downloads the exam schedule from AUEB
func download() ([]byte, bool) {
	var ret []byte

	resp, _ := http.Get(link)

	if resp.StatusCode == 200 {
		ret, _ = ioutil.ReadAll(resp.Body)
	} else {
		return exams, true
	}

	return ret, false
}

// configureLink Method that constructs link based on the date
func configureLink() {
	t := time.Now()
	date := t.Format("20060102")
	month := t.Month()
	link = fmt.Sprintf("https://aueb.gr/sites/default/files/aueb/%s_Exams_%s.xlsx", month, date)
}

// convertChars Method that converts english characters to greek in order to parse Schedule Master Room Name
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

// convertCharsInverse Method that converts greek characters to english in order to parse Schedule Master Room Name
func convertCharsInverse(room string) string {
	re := regexp.MustCompile("[0-9]+")

	if re.MatchString(room) {
		if strings.Contains(room, "Α") {
			room = strings.ReplaceAll(room, "Α", "a")
		}
		if strings.Contains(room, "Δ") {
			room = strings.ReplaceAll(room, "Δ", "d")
		}
		if strings.Contains(room, "Τ") {
			room = strings.ReplaceAll(room, "Τ", "t")
		}
		return room
	}

	return room
}