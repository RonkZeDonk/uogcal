package web

import (
	"fmt"
	"os"
	"strings"

	"github.com/RonkZeDonk/uogcal/pkg/calendar"
	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/aymerick/raymond"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func addContextHelpers(engine *handlebars.Engine) {
	engine.AddFunc("calendarContext", func(options *raymond.Options) raymond.SafeString {
		data, err := database.GetSectionsByUUID(options.ValueStr("id"))
		if err != nil {
			return raymond.SafeString(err.Error())
		}

		var dataMap []map[string]string

		for _, d := range data {
			dataMap = append(dataMap, map[string]string{
				"code":       d.Meeting.Code,
				"name":       d.Name,
				"startDate":  d.Meeting.StartDate.Format("Jan 02 2006"),
				"endDate":    d.Meeting.EndDate.Format("Jan 02 2006"),
				"courseType": d.Meeting.Type,
				"startTime":  d.Meeting.StartTime.Format("3:04 PM"),
				"endTime":    d.Meeting.EndTime.Format("3:04 PM"),
				"days":       fmt.Sprint(calendar.WeekdayFromIntArray(d.Meeting.MeetingDays)),
				"location":   d.Meeting.Location,
			})
		}

		return raymond.SafeString(options.FnWith(dataMap))
	})
}

func addHelpers(engine *handlebars.Engine) {
	engine.AddFunc("isLoggedIn", func(options *raymond.Options) bool {
		return options.ValueStr("user") != ""
	})

	addContextHelpers(engine)
}

func handleTemplates(c *fiber.Ctx) error {
	claims, _ := GetAuthClaims(c)

	route := c.Params("*")
	if route == "" {
		route = "index"
	}

	if strings.Split(route, "/")[0] == "layouts" {
		return notFoundMiddleware(c)
	}

	if _, err := os.Stat(fmt.Sprintf("views/%v", route)); err == nil {
		// Template not found
		if _, err := os.Stat(fmt.Sprintf("views/%v/index.hbs", route)); err == nil {
			// Sub-route found
			route = route + "/index"
		} else {
			return notFoundMiddleware(c)
		}
	}

	return c.Render(route, fiber.Map{
		"title": "UoG Course Calendar",
		"user":  claims.User,
		"id":    claims.Id,
		"isLoggedIn": func() bool {
			return claims.User != ""
		},
	}, "layouts/main")
}
