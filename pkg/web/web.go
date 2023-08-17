package web

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/handlebars/v2"
)

func notFoundMiddleware(c *fiber.Ctx) error {
	return c.Status(404).SendString("404 Not found")
}

func StartWeb() error {
	engine := handlebars.New("./views", ".hbs")
	if os.Getenv("ENV") == "development" {
		engine.Reload(true)
	}
	addHelpers(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Misc. Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(recover.New())
	app.Use(helmet.New())

	// Static files
	app.Static("/files", "./public")
	app.Static("/robots.txt", "robots.txt")

	// Routes
	app.Route("/", AuthRoutes)
	app.Route("/calendar", CalendarRoutes)
	app.Route("/upload", UploadRoutes)

	// Main templates route
	app.Get("/*", handleTemplates)

	// Last route left
	app.Use(notFoundMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "29944"
	}
	return app.Listen(fmt.Sprintf(":%v", port))
}
