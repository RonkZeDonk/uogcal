package web

import (
	"encoding/json"

	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func APIRoutes(r fiber.Router) {
	r.Route("/courses", func(cr fiber.Router) {
		cr.Get("", func(c *fiber.Ctx) error {
			clms, err := GetAuthClaims(c)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			arr, err := database.GetSectionsByUUID(clms.Id)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}

			data, err := json.Marshal(arr)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}

			return c.SendString(string(data))
		})
		cr.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.Status(fiber.ErrNotImplemented.Code).SendString(id + " not impl'd")
		})
	})
}
