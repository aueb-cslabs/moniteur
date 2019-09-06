package rest

import (
	"encoding/json"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CalendarGroup Defines the api paths for all the calendar information
func CalendarGroup(g *echo.Group) {
	g.GET("", calendarInfo)
}

// calendarInfo Method that GETs the state of the day
func calendarInfo(ec echo.Context) error {
	c := ec.(*types.Context)

	info := &types.Info{}

	info.IsWeekend = isWeekend()
	info.IsBreak = isBreak(c.Calendar.Breaks)
	info.IsExamsPeriod = isExams(c.Calendar.Exams)
	info.IsNormalPeriod = isNormal(c.Calendar.Semesters)

	return c.JSON(http.StatusOK, info)
}

// currentDate Method that returns the date in dd/mm/yyyy format.
func currentDate() string {
	now := time.Now()
	dateF := strconv.Itoa(now.Day()) + "/" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Year())

	return dateF
}

// checkDates Method that checks every if a date is between two date limits
func checkDates(start string, end string) bool {
	startA := strings.Split(start, "/")
	endA := strings.Split(end, "/")

	startY, startM, startD := convertDate(startA)
	endY, endM, endD := convertDate(endA)

	now := time.Now()
	startDate := time.Date(startY, startM, startD, 0, 0, 0, 0, now.Location())
	endDate := time.Date(endY, endM, endD, 0, 0, 0, 0, now.Location())

	return now.After(startDate) && now.Before(endDate)
}

// convertDate Method that breaks up a date
func convertDate(date []string) (int, time.Month, int) {
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	return year, time.Month(month), day
}

// isNormal Method that returns a boolean that indicates if the university is in normal period
func isNormal(semester types.Semester) bool {
	start := semester.Winter.Start
	end := semester.Winter.End

	res1 := checkDates(start, end)

	start = semester.Sprint.Start
	end = semester.Sprint.End

	res2 := checkDates(start, end)

	return (res1 || res2) && !isWeekend()
}

// isNormal Method that returns a boolean that indicates if the university is in exams period
func isExams(exams types.Exam) bool {
	start := exams.Winter.Start
	end := exams.Winter.End

	res1 := checkDates(start, end)

	start = exams.Sprint.Start
	end = exams.Sprint.End

	res2 := checkDates(start, end)

	start = exams.September.Start
	end = exams.September.End

	res3 := checkDates(start, end)

	return (res1 || res2 || res3) && !isWeekend()
}

// isNormal Method that returns a boolean that indicates if the university is in break
func isBreak(breaks types.Break) bool {
	start := breaks.Winter.Start
	end := breaks.Winter.End

	res1 := checkDates(start, end)

	start = breaks.Sprint.Start
	end = breaks.Sprint.End

	res2 := checkDates(start, end)

	res3 := false
	various := breaks.Various
	current := currentDate()

	res4 := saturateHolidAPI()

	for i := range various {
		if current == *various[i] {
			res3 = true
			break
		}
	}

	return res1 || res2 || res3 || res4
}

// isNormal Method that returns a boolean that indicates if it is weekend
func isWeekend() bool {
	now := time.Now()
	day := int(now.Weekday())

	if day == 0 || day == 6 {
		return true
	}
	return false
}

// saturateHolidAPI Method that saturates public Holid API for national holidays
func saturateHolidAPI() bool {
	resp, _ := http.Get("http://api.holid.co/v1/GR/Europe/Athens")
	bts, _ := ioutil.ReadAll(resp.Body)
	var holidAPI *types.HolidAPI
	_ = json.Unmarshal(bts, &holidAPI)

	holidays := holidAPI.Holidays

	if holidAPI.Success && len(holidays) != 0 {
		for i := range holidays {
			if strings.Contains(holidays[i].Name, "Solstice") ||
				strings.Contains(holidays[i].Name, "Equinox") ||
				strings.Contains(holidays[i].Name, "Armed") {
				return false
			}
		}
		return true
	}
	return false
}
