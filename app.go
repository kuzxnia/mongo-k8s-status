package main

import (
	"flag"
	"log"
	"mongo-k8s-status/database"
	"mongo-k8s-status/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Setup view engine
	engine := html.New("./views", ".html")

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
    Views: engine,
    ViewsLayout: "layouts/main",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())


	// Bind views
	app.Get("/", handlers.Home)
	app.Get("/k8s", handlers.K8s)
	app.Get("/clusters", handlers.Clusters)
	app.Get("/backups", handlers.Backups)

  // Bind routes
	// Create a /api/v1 endpoint
	// v1 := app.Group("/api/v1")

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
