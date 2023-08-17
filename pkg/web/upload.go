package web

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/RonkZeDonk/uogcal/pkg/collector"
	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

type UploadJson struct {
	Term    string
	Courses []struct {
		Title string
	}
}

func UploadRoutes(r fiber.Router) {
	r.Post("/courses", func(c *fiber.Ctx) error {
		claims, _ := GetAuthClaims(c)

		header, err := c.FormFile("courses")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		file, err := header.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		data, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		var result UploadJson
		err = json.Unmarshal(data, &result)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		verify, err := collector.GetVerificationTokenAndCookie()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		for _, section := range result.Courses {
			code := strings.Split(section.Title, ":")[0]
			if !database.CheckCourseExists(code) {
				course, meetings, err := collector.GetSectionData(result.Term, code, verify.Header, verify.Cookie)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
				}

				err = database.AddSection(course, meetings)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
				}
			}

			err = database.AddUserToSection(claims.Id, code)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/me/account")
	})
}
