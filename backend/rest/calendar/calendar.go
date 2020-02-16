package calendar

import (
	"encoding/json"
	"fmt"
	"github.com/aueb-cslabs/moniteur/backend/rest"
	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/aueb-cslabs/moniteur/backend/utils"
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
	g.GET("/full", calendarExt)
	g.POST("", authentication.Validate(postCalendar))
}

// calendarInfo Method that GETs the state of the day
func calendarInfo(ec echo.Context) error {
	info := &types.Info{}

	info.IsWeekend = isWeekend()
	info.IsBreak = isBreak(rest.Calendar.Breaks)
	info.IsExamsPeriod = isExams(rest.Calendar.Exams)
	info.IsNormalPeriod = isNormal(rest.Calendar.Semesters)

	return ec.JSON(http.StatusOK, info)
}

func calendarExt(e echo.Context) error {
	return e.JSON(http.StatusOK, rest.Calendar)
}

func postCalendar(e echo.Context) error {
	calendarIn := &types.Calendar{}

	err := e.Bind(calendarIn)
	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	rest.Calendar = calendarIn

	err = utils.UpdateCalendar(calendarIn)
	if err != nil {
		return e.NoContent(http.StatusInternalServerError)
	}

	return e.NoContent(http.StatusOK)
}

// currentDate Method that returns the date in dd/mm/yyyy format.
func currentDate() string {
	now := time.Now()
	dateF := strconv.Itoa(now.Day()) + "/" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Year())

	return dateF
}

// checkDates Method that checks every if a date is between two date limits
func checkDates(start string, end string) bool {
	startA := strings.Split(start, "-")
	endA := strings.Split(end, "-")

	startY, startM, startD := convertDate(startA)
	endY, endM, endD := convertDate(endA)

	now := time.Now()
	startDate := time.Date(startY, startM, startD, 0, 0, 0, 0, now.Location())
	endDate := time.Date(endY, endM, endD, 0, 0, 0, 0, now.Location()).Add(24 * time.Hour)

	return now.After(startDate) && now.Before(endDate)
}

// convertDate Method that breaks up a date
func convertDate(date []string) (int, time.Month, int) {
	day, _ := strconv.Atoi(date[2])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[0])

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
	resp, _ := http.Get(fmt.Sprintf("https://date.nager.at/api/v2/publicholidays/%d/GR", time.Now().Year()))
	bts, _ := ioutil.ReadAll(resp.Body)
	var holidays []types.Holiday
	_ = json.Unmarshal(bts, &holidays)

	if len(holidays) != 0 {
		now := time.Now()
		formatted := now.Format("2006-01-02")
		for i := range holidays {
			if formatted == holidays[i].Date {
				if strings.Contains(holidays[i].IntName, "Solstice") ||
					strings.Contains(holidays[i].IntName, "Equinox") ||
					strings.Contains(holidays[i].IntName, "Armed") {
					return false
				} else {
					return true
				}
			}
		}
	}
	return false
}
