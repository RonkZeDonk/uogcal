package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/olivere/vite"
)

const indexTemplate = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/icon.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>UoG Calendar</title>
	{{ .Vite.Tags }}
  </head>
  <body>
    <div id="root"></div>
  </body>
</html>`

func serveDev(r fiber.Router) {
	// serve /src/assets
	r.Get("/src/assets/*", func(c *fiber.Ctx) error {
		err := filesystem.SendFile(c, http.Dir("src/assets"), strings.TrimPrefix(c.Path(), "/src/assets"))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return nil
	})

	// serve public folder
	r.Use(func(c *fiber.Ctx) error {
		err := filesystem.SendFile(c, http.Dir("public"), c.Path())
		if err != nil {
			return c.Next()
		}
		return nil
	})

	// handle /
	r.Get("/*", func(c *fiber.Ctx) error {
		frag, err := vite.HTMLFragment(vite.Config{
			FS:      os.DirFS("."),
			IsDev:   true,
			ViteURL: "http://localhost:5173",
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			return err
		}

		tmpl, err := template.New("index").Parse(indexTemplate)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			return err
		}

		err = tmpl.Execute(c.Response().BodyWriter(), map[string]interface{}{
			"Vite": frag,
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			return err
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

		return nil
	})
}

func serveProd(static embed.FS) func(r fiber.Router) {
	return func(r fiber.Router) {
		distFS, err := fs.Sub(static, "dist")
		if err != nil {
			log.Panicln(err)
		}

		// serve assets
		r.Get("/assets/*", func(c *fiber.Ctx) error {
			err := filesystem.SendFile(c, http.FS(distFS), c.Path())
			if err != nil {
				return c.Status(fiber.StatusNotFound).SendString(err.Error())
			}
			return nil
		})

		// serve public files with .vite folder removed
		r.Use(func(c *fiber.Ctx) error {
			c.Path(strings.TrimPrefix(c.Path(), "/.vite"))
			err := filesystem.SendFile(c, http.FS(distFS), c.Path())
			if err != nil {
				return c.Next()
			}
			return nil
		})

		r.Get("/*", func(c *fiber.Ctx) error {
			frag, err := vite.HTMLFragment(vite.Config{
				FS: distFS,
			})
			if err != nil {
				c.Status(fiber.StatusInternalServerError).SendString(err.Error())
				return err
			}

			tmpl, err := template.New("index").Parse(indexTemplate)
			if err != nil {
				c.Status(fiber.StatusInternalServerError).SendString(err.Error())
				return err
			}

			err = tmpl.Execute(c.Response().BodyWriter(), map[string]interface{}{
				"Vite": frag,
			})
			if err != nil {
				c.Status(fiber.StatusInternalServerError).SendString(err.Error())
				return err
			}

			ext := filepath.Ext(c.Path())
			if ext == "" {
				ext = ".html"
			}
			c.Type(ext)

			return nil
		})
	}
}

// TODO is static assets really the best name for this?
func StaticAssets(static embed.FS) func(r fiber.Router) {
	if s, _ := os.LookupEnv("ENV"); s == "dev" {
		fmt.Println("Running dev server...")

		return serveDev
	} else {
		return serveProd(static)
	}
}
