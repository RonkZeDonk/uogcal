package database

import (
	"time"

	"github.com/google/uuid"
)

type CourseSection struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SectionMeeting struct {
	Code         string    `json:"code"`
	Type         string    `json:"type"`
	Created      time.Time `json:"created"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	MeetingDays  []int8    `json:"meetingDays"`
	Location     string    `json:"location"`
	LastModified time.Time `json:"lastModified"`
	UpdateCount  int       `json:"updateCount"`
}

type CourseSectionJoin struct {
	Name    string         `json:"name"`
	Meeting SectionMeeting `json:"sectionMeeting"`
}

type UOGCalUser struct {
	UID         uuid.UUID `json:"uuid"`
	DisplayName string    `json:"displayName"`
	Password    string    `json:"password"`
}
