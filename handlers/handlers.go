package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Home renders the home view
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

// About renders the about view
func K8s(c *fiber.Ctx) error {
	return c.Render("k8s", nil)
}
func Clusters(c *fiber.Ctx) error {
	return c.Render("clusters", nil)
}
func Backups(c *fiber.Ctx) error {
	return c.Render("backups", nil)
}

// NoutFound renders the 404 view
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", nil)
}

// // NotFound returns custom 404 page
// func NotFound(c *fiber.Ctx) error {
// 	return c.Status(404).SendFile("./static/private/404.html")
// }
