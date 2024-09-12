package web

import (
	"embed"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartWeb(static embed.FS) error {
	app := fiber.New(fiber.Config{})

	// Misc. Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover.New())
	// TODO temporarily disabled for Google OAuth to work.
	//  - CSFR is the prime suspect for breaking OAuth
	// app.Use(helmet.New())

	// Routes
	app.Route("/", AuthRoutes)
	app.Route("/calendar", CalendarRoutes)
	app.Route("/upload", UploadRoutes)
	app.Route("/api", APIRoutes)

	// Static routes
	app.Use("/", StaticAssets(static))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusInternalServerError)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "29944"
	}
	return app.Listen(fmt.Sprintf(":%v", port))
}
