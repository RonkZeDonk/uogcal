package web

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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

	// Vite routes
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(static),
		PathPrefix:   "/dist",
		NotFoundFile: "/dist/index.html",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "29944"
	}
	return app.Listen(fmt.Sprintf(":%v", port))
}
