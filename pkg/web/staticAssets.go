package web

import (
	"embed"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func StaticAssets(static embed.FS) func(ctx *fiber.Ctx) error {
	if s, _ := os.LookupEnv("ENV"); s == "dev" {
		// TODO
		return func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
	} else {
		return filesystem.New(filesystem.Config{
			Root:         http.FS(static),
			PathPrefix:   "/dist",
			NotFoundFile: "/dist/index.html",
		})
	}
}
