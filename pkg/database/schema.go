package database

import (
	"time"

	"github.com/google/uuid"
)

type CourseSection struct {
	Code string
	Name string
}

type SectionMeeting struct {
	Code         string
	Type         string
	Created      time.Time
	StartDate    time.Time
	EndDate      time.Time
	StartTime    time.Time
	EndTime      time.Time
	MeetingDays  []int8
	Location     string
	LastModified time.Time
	UpdateCount  int
}

type CourseSectionJoin struct {
	Name    string
	Meeting SectionMeeting
}

type UOGCalUser struct {
	UID         uuid.UUID
	DisplayName string
	Password    string
}
