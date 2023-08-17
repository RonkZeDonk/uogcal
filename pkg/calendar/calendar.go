package calendar

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type icsEvent struct {
	title       string
	description string
	location    string

	created       time.Time
	lastModified  time.Time
	updateCount   int
	startDate     time.Time
	endDate       time.Time
	repeatingDays []time.Weekday

	startTime time.Time
	endTime   time.Time
}

type ics struct {
	calendarName string

	events *list.List
}

func (c *ics) AddEvent(
	title string,
	description string,
	location string,
	startDate time.Time,
	endDate time.Time,
	days []time.Weekday,
	startTime time.Time,
	endTime time.Time,
	created time.Time,
	lastModified time.Time,
	updateCount int,
) {
	c.events.PushBack(icsEvent{
		title:         title,
		description:   description,
		location:      location,
		created:       created,
		lastModified:  lastModified,
		updateCount:   updateCount,
		startDate:     startDate,
		endDate:       endDate,
		repeatingDays: days,
		startTime:     startTime,
		endTime:       endTime,
	})
}

func WeekdayFromIntArray(arr []int8) []time.Weekday {
	lut := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	weekdays := []time.Weekday{}

	for _, day := range arr {
		weekdays = append(weekdays, lut[day])
	}

	return weekdays
}

func timeToString(t time.Time) string {
	return t.Format("20060102T150405")
}
func joinDateWithClockTime(d time.Time, t time.Time) time.Time {
	return time.Date(
		d.Year(),
		d.Month(),
		d.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.UTC,
	)
}
func weekdayArrayToICSString(arr []time.Weekday) string {
	res := []string{}

	lut := []string{"SU", "MO", "TU", "WE", "TH", "FR", "SA"}

	for _, day := range arr {
		res = append(res, lut[day])
	}

	return strings.Join(res, ",")
}
func getNextNthWeekday(date time.Time, weekdays []time.Weekday) time.Time {
	if len(weekdays) == 0 {
		// TODO get rid of this panic
		panic("Not enough weekdays")
	}
	for {
		for _, weekday := range weekdays {
			if date.Weekday() == weekday {
				return date
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func buildEvent(evt icsEvent) string {
	return fmt.Sprintf(
		`BEGIN:VEVENT
        DTSTART;TZID=America/New_York:%v
        DTEND;TZID=America/New_York:%v
        RRULE:FREQ=WEEKLY;WKST=SU;UNTIL=%v;BYDAY=%v
        UID:%v
        DTSTAMP;TZID=America/New_York:%v
        LAST-MODIFIED;TZID=America/New_York:%v
        SEQUENCE:%v
        SUMMARY:%v
        DESCRIPTION:%v
        LOCATION:%v
        STATUS:CONFIRMED
        TRANSP:OPAQUE
        END:VEVENT`,

		// Start time (EST/DST)
		timeToString(getNextNthWeekday(
			joinDateWithClockTime(evt.startDate, evt.startTime),
			evt.repeatingDays,
		)),
		// Duration in minutes
		timeToString(getNextNthWeekday(
			joinDateWithClockTime(evt.startDate, evt.endTime),
			evt.repeatingDays,
		)),

		// Last day (EST/DST)
		timeToString(evt.endDate),
		// Repeating days
		weekdayArrayToICSString(evt.repeatingDays),

		// Unique event ID
		fmt.Sprintf("%v.%v.%v", evt.created.Format(time.DateOnly), evt.created.Format(time.TimeOnly), evt.created.Nanosecond()),
		// Event created time and date (EST/DST)
		timeToString(evt.created),
		// Last modified (EST/DST)
		timeToString(evt.lastModified),
		// Number of updates
		evt.updateCount,

		// Title
		evt.title,
		// Description
		evt.description,
		// Location
		evt.location,
	)
}

func (c *ics) generateEvents() string {
	res := ""

	for e := c.events.Front(); e != nil; e = e.Next() {
		evt := buildEvent(e.Value.(icsEvent))

		res = strings.Join([]string{res, evt}, "\n")
	}

	return res
}

func (c *ics) GenerateCalendar() string {
	events := c.generateEvents()

	cal := fmt.Sprintf(
		`BEGIN:VCALENDAR
        PRODID:-//RonkZD//Course Calendar v0.0.1//EN
        VERSION:2.0
        CALSCALE:GREGORIAN
        METHOD:PUBLISH
        X-WR-CALNAME:%v
        X-WR-TIMEZONE:America/New_York
        BEGIN:VTIMEZONE
        TZID:America/New_York
        X-LIC-LOCATION:America/New_York
        BEGIN:DAYLIGHT
        TZOFFSETFROM:-0500
        TZOFFSETTO:-0400
        TZNAME:EDT
        DTSTART:19700308T020000
        RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=2SU
        END:DAYLIGHT
        BEGIN:STANDARD
        TZOFFSETFROM:-0400
        TZOFFSETTO:-0500
        TZNAME:EST
        DTSTART:19701101T020000
        RRULE:FREQ=YEARLY;BYMONTH=11;BYDAY=1SU
        END:STANDARD
        END:VTIMEZONE
        %v
        END:VCALENDAR`,
		c.calendarName,
		events,
	)

	// The following code _shouldn't_ ever error

	// Remove all spaces (from string indentation) and empty lines
	rg, err := regexp.Compile(`(?m:^\s+)`)
	if err != nil {
		return ""
	}
	cal = rg.ReplaceAllString(cal, "")
	// Convert to CRLF line endings as required by standard
	rg, err = regexp.Compile(`\r?\n`)
	if err != nil {
		return ""
	}
	cal = rg.ReplaceAllString(cal, "\r\n")

	return cal
}

func NewCalendar(title string) ics {
	return ics{
		calendarName: title,
		events:       list.New(),
	}
}
