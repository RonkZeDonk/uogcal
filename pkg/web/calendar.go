package web

import (
	"fmt"
	"time"

	"github.com/RonkZeDonk/uogcal/pkg/calendar"
	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func CalendarRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		claims, err := GetAuthClaims(c)
		if err != nil {
			c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
			return c.Redirect("/calendar/" + claims.Id)
		} else {
			// Demo event
			cal := calendar.NewCalendar("UoG Calendar Demo")
			now := time.Now()
			cal.AddEvent(
				"Demo Course",
				"uogcal.ronkzd.xyz",
				"Guelph, ON",
				now,
				now.Add(time.Hour),
				[]time.Weekday{now.Weekday()},
				now,
				now,
				now,
				now,
				0,
			)

			c.Type("text/calendar")
			c.Attachment("courses.ics")
			return c.SendString(cal.GenerateCalendar())
		}
	})

	r.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		data, err := database.GetSectionsByUUID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(
				"Error (p1)",
			)
		}

		cal := calendar.NewCalendar("UoGuelph Courses")

		for _, d := range data {
			cal.AddEvent(
				fmt.Sprintf("(%v) %v", d.Meeting.Type, d.Name),
				"Event managed by uogcal (https://uogcal.ronkzd.xyz/)",
				d.Meeting.Location,
				d.Meeting.StartDate,
				d.Meeting.EndDate,
				calendar.WeekdayFromIntArray(d.Meeting.MeetingDays),
				d.Meeting.StartTime,
				d.Meeting.EndTime,
				d.Meeting.Created,
				d.Meeting.LastModified,
				d.Meeting.UpdateCount,
			)
		}

		c.Type("text/calendar")
		c.Attachment("courses.ics")
		return c.SendString(cal.GenerateCalendar())
	})
}
