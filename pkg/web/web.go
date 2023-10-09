package web

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartWeb() error {
	app := fiber.New(fiber.Config{})

	// Misc. Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover.New())
	app.Use(helmet.New())

	// Static files
	app.Static("/", "./dist")

	// Routes
	app.Route("/", AuthRoutes)
	app.Route("/calendar", CalendarRoutes)
	app.Route("/upload", UploadRoutes)
	app.Route("/api", APIRoutes)

	// Vite routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "29944"
	}
	return app.Listen(fmt.Sprintf(":%v", port))
}
